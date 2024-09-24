package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"strings"
	"github.com/golang-jwt/jwt/v5"
	"github.com/go-chi/httprate"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
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
	//16. implement CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
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

	//mount grocery route to the main function
	r.Mount("/groceries/v1", GroceryRoutesV1())
	r.Mount("/groceries/v2", GroceryRoutesV2())

	http.ListenAndServe(":3000", r)
	
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