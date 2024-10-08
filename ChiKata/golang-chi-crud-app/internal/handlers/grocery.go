package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/chloejepson16/golang-chi-crud-app/internal/models"

	"github.com/go-chi/chi/v5"
)

//contains functions to perform crud operations
type GroceryHandler struct{
	DB *sql.DB
}

func (g GroceryHandler) ListGroceries( w http.ResponseWriter, r *http.Request) {
	//8. implement error handling
	err:= json.NewEncoder(w).Encode(models.ListGroceries())
	if err != nil{
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

//3. create a route that accepts query params
func (g GroceryHandler) GetGroceries( w http.ResponseWriter, r *http.Request) {
	//6. implement route params with chi to retrieve specific data
	id:= chi.URLParam(r, "id")
	grocery:= models.GetGroceries(id)
	if grocery == nil{
		http.Error(w, "No grocery item found", http.StatusNotFound)
	}
	err:= json.NewEncoder(w).Encode(grocery)
	if err != nil{
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
//5. create a route that accepts post requests
func (g GroceryHandler) CreateGroceries( w http.ResponseWriter, r *http.Request) {
	var grocery models.GroceryItem
	err:= json.NewDecoder(r.Body).Decode(&grocery)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.AddGrocery(grocery)
	err= json.NewEncoder(w).Encode(grocery)
	if err != nil{
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
func (g GroceryHandler) UpdateGroceries( w http.ResponseWriter, r *http.Request) {
	id:= chi.URLParam(r, "id")
	var grocery models.GroceryItem
	err:= json.NewDecoder(r.Body).Decode(&grocery)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedGrocery:= models.UpdateGrocery(id, grocery)
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
	grocery:= models.DeleteGrocery(id)
	if grocery == nil{
		http.Error(w, "Grocery not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (g GroceryHandler) AddGroceryToDB( w http.ResponseWriter, r *http.Request){
	var item models.GroceryItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Fatalf("Failed to open a DB connection on add to grocerydb: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if g.DB == nil {
		http.Error(w, "Database connection is not available", http.StatusInternalServerError)
		return
	}
	if err := models.AddGroceryItemToDB(g.DB, item); err != nil {
		fmt.Println("failed at models.add grocery to db")
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (g GroceryHandler) GetGroceriesFromDB( w http.ResponseWriter, r *http.Request){
	if g.DB == nil {
		http.Error(w, "Database connection is not available", http.StatusInternalServerError)
		return
	}

	groceriesJSON, err := models.GetGroceryListFromDB(g.DB)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
    w.Write(groceriesJSON)
}

//13. create a route using chi that fetches data from an extrenal api
func (g GroceryHandler) GetJellyBeans(w http.ResponseWriter, r *http.Request){
	flavorName:= chi.URLParam(r, "flavorName")
	jellyBean, err:= models.GetJellyBeans(flavorName)
	if err != nil{
		fmt.Println(err)
		http.Error(w, "JellyBean was not found", http.StatusNotFound)
		return
	}
	if jellyBean == "" {
		fmt.Println(err)
        http.Error(w, "JellyBean was not found", http.StatusNotFound)
        return
    }
	w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(jellyBean))
}

func (g GroceryHandler) GetV2 (w http.ResponseWriter, r *http.Request){
	//18. implementing a shutdown based on timeout of request
	//create context that cancels the operation after 5 seconds
	ctx, cancel:= context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	//create a channel for response from data
	responseCh:= make(chan string)
	errCh:= make(chan error)

	//launch goroutine to handle the request and send result or error
	go func(){
		response, err:= models.GetJellyBeans("7up")
		if err != nil{
			errCh <- err
			return
		}
		responseCh <- response
	}()
	
	//listen to response from goroutine 
	select{
		case response := <-responseCh:
			if response == "" {
				http.Error(w, "JJellyBean was not found (endpoit to read in chunks)", http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			//17. write data in chunks
			chunkSize:= 50
			for i:= 0; i< len(response); i+= chunkSize{
				end:= i+ chunkSize
				if end > len(response){
					end= len(response)
				}
		
				chunk:= response[i:end]
				_, err := w.Write([]byte(fmt.Sprintf("%q", chunk)))
				if err != nil{
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
		
				w.(http.Flusher).Flush()
			}
			w.Write([]byte("end of jelly bean stream response"))
			
		case err:= <-errCh:
			http.Error(w, fmt.Sprintf("Error retrieving jelly beans: %v", err), http.StatusBadRequest)
		case <-ctx.Done():
			http.Error(w, "Request timed out", http.StatusGatewayTimeout)
	}

}

//TODO: separate this out into models.go
//file upload works with this curl command : curl -F "file=@/Users/chloejepson/Documents/Gokata/textExample.txt" http://localhost:3000/groceries/fileUpload
func (g GroceryHandler) UploadFile(w http.ResponseWriter, r *http.Request){
	err:= r.ParseMultipartForm(10 << 20)
	if err != nil{
		http.Error(w, "cannot parse form", http.StatusBadRequest)
		return
	}

	file, handler, err:= r.FormFile("file")
	if err != nil{
		http.Error(w, "unable to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	savePath:= filepath.Join("/Users/chloejepson/Documents/Gokata/Chi Kata", handler.Filename)

	dest, err:= os.Create(savePath)
	if err != nil{
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	defer dest.Close()

	_, err= dest.ReadFrom(file)
	if err != nil{
		http.Error(w, "issue saving file", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("File uploaded successfully: %s\n", handler.Filename)))
}