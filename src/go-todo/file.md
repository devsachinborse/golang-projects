package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	id         int       `json:"id"`
	title      string    `json:"title"`
	status     bool      `json:"status"`
	created_at time.Time `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
}

var todo = []Todo{
	{id: 1, title: "learning golang", status: false, created_at: time.Now(), updated_at: time.Now()},
}

func gettodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("contenet-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func main() {

	address := ":8080"

	router := mux.NewRouter()

	router.HandleFunc("/todos", gettodo).Methods("GET")

	fmt.Printf("server listning...")
	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Fatal(err)
	}

}
