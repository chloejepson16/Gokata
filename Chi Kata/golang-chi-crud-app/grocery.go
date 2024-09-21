package main

import(
"net/http"
"encoding/json"

"github.com/go-chi/chi/v5") 

//contains functions to perform crud operations
type GroceryHandler struct{
}

func (g GroceryHandler) ListGroceries( w http.ResponseWriter, r *http.Request) {
	err:= json.NewEncoder(w).Encode(listGroceries())
	if err != nil{
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (g GroceryHandler) GetGroceries( w http.ResponseWriter, r *http.Request) {
	id:= chi.URLParam(r, "id")
	grocery:= getGroceries(id)
	if grocery == nil{
		http.Error(w, "No grocery item found", http.StatusNotFound)
	}
	err:= json.NewEncoder(w).Encode(grocery)
	if err != nil{
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (g GroceryHandler) CreateGroceries( w http.ResponseWriter, r *http.Request) {
	var grocery GroceryItem
	err:= json.NewDecoder(r.Body).Decode(&grocery)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	addGrocery(grocery)
	err= json.NewEncoder(w).Encode(grocery)
	if err != nil{
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
func (g GroceryHandler) UpdateGroceries( w http.ResponseWriter, r *http.Request) {
	id:= chi.URLParam(r, "id")
	var grocery GroceryItem
	err:= json.NewDecoder(r.Body).Decode(&grocery)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedGrocery:= updateGrocery(id, grocery)
	if updatedGrocery == nil{
		http.Error(w, "Grocery not found", http.StatusNotFound)
		return
	}
	err= json.NewEncoder(w).Encode(updatedGrocery)
	if err != nil{
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
func (g GroceryHandler) DeleteGroceries( w http.ResponseWriter, r *http.Request) {
	id:= chi.URLParam(r, "id")
	grocery:= deleteGrocery(id)
	if grocery == nil{
		http.Error(w, "Grocery not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}