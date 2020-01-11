package main

import (
	"gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	get := r.PathPrefix("/get").Subrouter()
	get.HandleFunc("/{id}", GetByID)
	get.HandleFunc("/name/{name}", GetByName)

	post := r.PathPrefix("/post").Subrouter()
	post.HandleFunc("/new", AddNew).Methods(http.MethodPost)
	http.ListenAndServe(":8080", r)

}
