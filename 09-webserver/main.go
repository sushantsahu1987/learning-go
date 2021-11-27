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
	json.NewEncoder(w).Encode(posts)
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var post Post

	for _, p := range posts {
		if p.ID == params["id"] {
			post = p
		}
	}

	json.NewEncoder(w).Encode(post)
}

// TODO
func updatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post Post
	var updatedPost Post

	err := json.NewDecoder(r.Body).Decode(&updatedPost)
	for _, p := range posts {
		if p.ID == updatedPost.ID {
			post = p
			break
		}
	}

	post.Title = updatedPost.Title
	post.Author = updatedPost.Author

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "an error has occurred"},
		)
	}

	json.NewEncoder(w).Encode(posts)

}

// TODO
func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hello world from DELETE POST"},
	)

}

func addPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post Post
	json.NewDecoder(r.Body).Decode(&post)
	posts = append(posts, post)
	json.NewEncoder(w).Encode(posts)
}

type Post struct {
	ID     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
}

var posts = []Post{
	{
		ID:     "1",
		Title:  "Learning Go",
		Author: "Sushant Sahu",
	},
	{
		ID:     "2",
		Title:  "Fashion By Day",
		Author: "Khushboo Sewak",
	},
	{
		ID:     "3",
		Title:  "Fashion By Night",
		Author: "Khushboo Sewak",
	},
	{
		ID:     "4",
		Title:  "Learning Spark",
		Author: "Sushant Sahu",
	},
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
