package main

//import pakcages
import (
	"fmt"
	"log"
	"net/http"
)

// handler function to handle incomming HTTP requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/helloworld" { // Check the request URL path
		http.Error(w, "404 not found", http.StatusNotFound) // If the request URL path is not "/", return a 404 Not Found error
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound) // Check the request Method
		return
	}
	fmt.Fprintf(w, "Hello world") // Write the response back to the client
}

// main function
func main() {

	// Declare HTTP server on port 8080 by default it's server running on port 8080
	port := ":8000"

	// Register the handler function for the URL ("/hellworld")
	http.HandleFunc("/helloworld", helloHandler)

	fmt.Println("🚀Server is running on http://localhost:8000")
	err := http.ListenAndServe(port, nil) // The second parameter is an HTTP handler, which is nil here (uses DefaultServeMux)
	if err != nil {                       // If an error occurs, it will be logged by the Fatal function and the program will exit
		log.Fatal(err)
	}

}

//to build the project
// > gomod init
// > go build
// to run
// > go run main.go
