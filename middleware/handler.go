package middleware

import "net/http"

import "github.com/gorilla/mux"
func Handle() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Breakfast recipies", ))
	})
	return r
}