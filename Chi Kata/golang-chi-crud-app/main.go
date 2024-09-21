package main

import(
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	//1. create a simple http server that responds with hello world
    r := chi.NewRouter()
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
func GroceryRoutes() chi.Router{
	r:= chi.NewRouter()
	groceryHandler:= GroceryHandler{}
	r.Get("/", groceryHandler.ListGroceries)
	r.Post("/", groceryHandler.CreateGroceries)
	r.Get("/{id}", groceryHandler.GetGroceries)
	r.Put("/{id}", groceryHandler.UpdateGroceries)
	r.Delete("/{id}", groceryHandler.DeleteGroceries)
	return r
}