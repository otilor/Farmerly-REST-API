package api

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(getCategories())

}
func Users(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(getUsers())
}