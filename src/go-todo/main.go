package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

var todos []Todo

//get todos
func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("contenet-Type", "application/json")

	todos, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(todos)

}

//create todo

func createTodo(w http.ResponseWriter, r *http.Request) {

	var newtodo Todo
	_ = json.NewDecoder(r.Body).Decode(&newtodo)
	newtodo.Created_at = time.Now()
	newtodo.Updated_at = time.Now()
	newtodo.Id = len(todos) + 1
	todos = append(todos, newtodo)
	json.NewEncoder(w).Encode(newtodo)

}

//get todo by id

func getTodoById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Contenet-Type", "application-type")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"]) //use this if the id type is int

	for _, item := range todos {
		if item.Id == id {

			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})

}

//update todo

func updateTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var updatedTodo Todo
	_ = json.NewDecoder(r.Body).Decode(&updatedTodo)
	updatedTodo.Id = id
	updatedTodo.Updated_at = time.Now()

	for i, item := range todos {
		if item.Id == id {
			todos[i] = updatedTodo
			break
		}

	}
	json.NewEncoder(w).Encode(updatedTodo)

}

//delete fuction

func deleteTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range todos {
		if item.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

func main() {

	todos = append(todos, Todo{Id: 1, Title: "learning golang", Status: "pending", Created_at: time.Now(), Updated_at: time.Now()})

	address := ":8080"

	router := mux.NewRouter()

	router.HandleFunc("/todos", getTodo).Methods("GET")
	router.HandleFunc("/todos", createTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", getTodoById).Methods(("GET"))
	router.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	fmt.Printf("server listning...")
	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Fatal(err)
	}

}
