package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var tasks = initData()

func getTaskHandler(w http.ResponseWriter, r *http.Request) {

	path := mux.Vars(r)
	w.Header().Add("Content-Type", "application/json")
	id, err := strconv.Atoi(path["id"])
	if err != nil {
		resp := map[string]string{
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
	}

	idx := sort.Search(len(tasks), func(i int) bool { return tasks[i].ID >= id })
	if idx < len(tasks) && tasks[idx].ID == id {
		json.NewEncoder(w).Encode(tasks[idx])
	} else {
		resp := map[string]string{
			"id": path["id"] + " not found",
		}
		json.NewEncoder(w).Encode(resp)
	}

}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	w.Header().Add("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&task)
	task.ID = len(tasks) + 1
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	r.Header.Add("Content-Type", "application/json")
	idx := sort.Search(len(tasks), func(i int) bool { return tasks[i].ID >= task.ID })
	if idx < len(tasks) && tasks[idx].ID == task.ID {
		tasks[idx].Description = task.Description
		tasks[idx].Status = task.Status
		json.NewEncoder(w).Encode(tasks[idx])
	} else {
		resp := map[string]string{
			"task_id": strconv.Itoa(task.ID) + " not found",
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)
	w.Header().Add("Content-Type", "application/json")
	id, err := strconv.Atoi(path["id"])
	if err != nil {
		resp := map[string]string{
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
	}
	idx := sort.Search(len(tasks), func(i int) bool { return tasks[i].ID >= id })
	if idx < len(tasks) && tasks[idx].ID == id {
		temp := tasks[idx]
		tasks = append(tasks[:idx], tasks[idx+1:]...)
		json.NewEncoder(w).Encode(temp)
	} else {
		resp := map[string]string{
			"task_id": path["id"] + " not found",
		}
		json.NewEncoder(w).Encode(resp)
	}

}

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func initData() []Task {
	tasks := []Task{
		{ID: 1, Description: "Walk my dog", Status: true},
		{ID: 2, Description: "Walk my cat", Status: false},
		{ID: 3, Description: "Grocery shopping", Status: true},
		{ID: 4, Description: "Clean my room", Status: false},
		{ID: 5, Description: "Clean my car", Status: true},
		{ID: 6, Description: "Clean my garage", Status: false},
	}
	return tasks
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tasks/{id}", getTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", getTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", addTaskHandler).Methods("POST")
	r.HandleFunc("/tasks", updateTaskHandler).Methods("PUT")
	r.HandleFunc("/tasks/{id}", deleteTaskHandler).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server started on port 127.0.0.1:8080")
	log.Fatal(srv.ListenAndServe())
}
