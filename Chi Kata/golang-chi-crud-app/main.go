package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"time"
	"os"
	"path/filepath"
	"log"
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"strings"
	"github.com/golang-jwt/jwt/v5"
	"github.com/go-chi/httprate"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r*http.Request) bool{
		return true
	},
	ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
	//1. create a simple http server that responds with hello world
    r := chi.NewRouter()
	//4. implement chi middleware
    r.Use(middleware.Logger)
	//23. compress
	r.Use(middleware.Compress(5))
	//16. implement CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	//23. connect a SQL Db
	//export all env variables, un and pw are 'postgres', port is 5432, host is go-kata-db.czics24ggkzl.us-east-1.rds.amazonaws.com
	dbUser:= os.Getenv("DB_USER")
	dbPassword:= os.Getenv("DB_PASSWORD")
	dbName:= os.Getenv("DB_NAME")
	dbHost:= os.Getenv("DB_HOST")
	dbPort:= os.Getenv("DB_PORT")

	//connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err:= sql.Open("postgres", dsn)
	if err != nil{
		log.Fatalf("Failed to open a DB connection: %v", err)
	}
	defer db.Close()

	err= db.Ping()
	if err != nil{
		log.Fatalf("Failed to open a DB connection: %v", err)
	}

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World"))
    })

	//2. create a route that returns a json response
	r.Get("/json", func(w http.ResponseWriter, r *http.Request){
		response:= map[string]string{
			"message": "this is the json response instead of Hello, World",
		}
		w.Header().Set("Content-Type", "applicaiton/json")
		json.NewEncoder(w).Encode(response)
	})
	//adding web socket
	r.Get("/ws", webSocketEndpoint)
	//22. route for serving static files
	workDir, _:= os.Getwd()
	filesDir:= http.Dir(filepath.Join(workDir,"resources"))
	FileServer(r, "/resources", filesDir)

	//mount grocery route to the main function
	r.Mount("/groceries/v1", GroceryRoutesV1())
	r.Mount("/groceries/v2", GroceryRoutesV2())

	http.ListenAndServe(":3000", r)
	
}

func FileServer(r chi.Router, path string, root http.FileSystem){
	if strings.ContainsAny(path, "{}*") {
		panic("cannot use URL params")
	}

	if path != "/" && path[len(path)-1] != '/'{
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path+="/"
	}
	//include all file types within resources directory
	path +="*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request){
		rctx:= chi.RouteContext(r.Context())
		pathPrefix:= strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs:= http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func JWTAuth(next http.Handler) http.Handler{
	var jwtKey= []byte("qwertyuiopasdfghjklzxcvbnm123456")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader:= r.Header.Get("Authorization")
		if authHeader == ""{
			http.Error(w, "missing auth header", http.StatusUnauthorized)
		}

		tokenString:= strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader{
			http.Error(w, "malformed auth header",  http.StatusUnauthorized)
			return
		}

		token, err:= jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
			if _, ok:= token.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil, http.ErrAbortHandler
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid{
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

//19. create a websocket endpoint for real time comms
func webSocketEndpoint(w http.ResponseWriter, r *http.Request){
	ws, err:= upgrader.Upgrade(w, r, nil)
	if err != nil{
		http.Error(w, "can't upgrade connection", http.StatusInternalServerError)
		return
	}
	defer ws.Close()

	//listen for messages from the websocket
	for{
		messageType, message, err:= ws.ReadMessage()
		if err!= nil{
			fmt.Println("Read error: ", err)
			break
		}
		fmt.Println(message)
		err= ws.WriteMessage(messageType, message)
		if err != nil{
			fmt.Println(err)
			break
		}
	}
}


//mounting grocery handler
//7. create an endpoint to perform crud operations on mock data
func GroceryRoutesV1() chi.Router{
	r:= chi.NewRouter()
	groceryHandler:= GroceryHandler{}
	//9. add authentification to one of the endpoints
	//12. add rate limiting to one endpoint
	r.With(httprate.Limit(2, 1*time.Minute)).Get("/groceries", groceryHandler.ListGroceries)
	r.With(JWTAuth).Get("/", groceryHandler.ListGroceries)
	r.Post("/", groceryHandler.CreateGroceries)
	r.Get("/{id}", groceryHandler.GetGroceries)
	r.Put("/{id}", groceryHandler.UpdateGroceries)
	r.Delete("/{id}", groceryHandler.DeleteGroceries)
	r.Post("/fileUpload", groceryHandler.UploadFile)
	//13. create a rout that fetches data from another api: http://localhost:3000/groceries/jellybeans/7up
	r.Get("/jellybeans/{flavorName}", groceryHandler.GetJellyBeans)
	return r
}
//15. create a route that supports versioning
func GroceryRoutesV2() chi.Router{
	r:= chi.NewRouter()
	groceryHandler:= GroceryHandler{}
	r.With(httprate.Limit(2, 1*time.Minute)).Get("/groceries", groceryHandler.ListGroceries)
	r.Get("/", groceryHandler.GetV2)
	return r
}