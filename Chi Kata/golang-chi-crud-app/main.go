package main

import(
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"strings"
	"github.com/golang-jwt/jwt/v5"
)


var jwtKey= []byte("qwertyuiopasdfghjklzxcvbnm123456")
func JWTAuth(next http.Handler) http.Handler{
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
func main() {
	//1. create a simple http server that responds with hello world
    r := chi.NewRouter()
	//4. implement chi middleware
    r.Use(middleware.Logger)
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

	//mount grocery route to the main function
	r.Mount("/groceries", GroceryRoutes())

	http.ListenAndServe(":3000", r)
}

//mounting grocery handler
//7. create an endpoint to perform crud operations on mock data
func GroceryRoutes() chi.Router{
	r:= chi.NewRouter()
	groceryHandler:= GroceryHandler{}
	//9. add authentification to one of the endpoints
	r.With(JWTAuth).Get("/", groceryHandler.ListGroceries)
	r.Post("/", groceryHandler.CreateGroceries)
	r.Get("/{id}", groceryHandler.GetGroceries)
	r.Put("/{id}", groceryHandler.UpdateGroceries)
	r.Delete("/{id}", groceryHandler.DeleteGroceries)
	return r
}