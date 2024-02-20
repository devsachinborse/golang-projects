package main

//import pakcages
import (
	"fmt"
	"log"
	"net/http"
)

// handler function to handle incomming HTTP requests

func helloHandler (w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "Method not found" , http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello...!") 
}


func formHandler(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succesffully \n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w,"Name =  %s\n", name )
	fmt.Fprintf(w,"Email =  %s\n", email )
}


func main() {

	port := ":8000"

	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("server is running on port:8000")
	err:= http.ListenAndServe(port,nil)
	if err != nil {
		log.Fatal(err)
	}
}











// main function
// func main() {

// 	// Declare HTTP server on port 8080 by default it's server running on port 8080

// 	port := ":8000"

// 	// Serve static files from the "static" directory

// 	fileServer := http.FileServer(http.Dir("./static"))

// 	// Register the handler function for the URL ("/hellworld")

// 	http.Handler("/" , fileServer)
// 	http.HandleFunc("/welcome", welcomeHandler)

// 	fmt.Println("🚀Server is running on http://localhost:8000")
// 	err := http.ListenAndServe(port, nil) // The second parameter is an HTTP handler, which is nil here (uses DefaultServeMux)
// 	if err != nil {                       // If an error occurs, it will be logged by the Fatal function and the program will exit
// 		log.Fatal(err)
// 	}

// }

//to build the project
// > gomod init
// > go build
// to run
// > go run main.go
