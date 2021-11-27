package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hello world from GET POSTS"},
	)
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hello world from GET POST"},
	)
}

func updatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hello world from PUT POST"},
	)
}

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hello world from DELETE POST"},
	)

}

func addPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hello world from POST POST"},
	)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/posts", getPostsHandler).Methods("GET")
	r.HandleFunc("/posts/{id}", getPostHandler).Methods("GET")
	r.HandleFunc("/posts", addPostHandler).Methods("POST")
	r.HandleFunc("/posts", updatePostHandler).Methods("PUT")
	r.HandleFunc("/posts/{id}", deletePostHandler).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server started at localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
