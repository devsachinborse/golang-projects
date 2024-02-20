package main

/*pakages:
fmt - to print outputs
log - to print logs
net/http - to create a server
math/rand - to create for user id
encoding/json - encod data into json
strconv - to convert integer to string
github.com/gorilla/mux - HTTP router and URL matcher
*/
import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
)

//defined struct
type Movie struct{
	ID				string 			`json:"id"`
	Isbn 			string			`json:"isbn"`
	Title 			string 			`json:"title"`
	Director 		*Director 		`json:"director"`
}

type Director struct{
	FirstName			string 			`json:"firstname"`
	LastName 			string			`json:"lastname"`
}

//defined slice of Movie
var movies[]Movie


//Handler functions

//create movie handler
func createMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}


//getMovies handler OK
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

//delete hangler
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies{
		
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

//get movie handler
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	for _, item := range movies   {                               //use _ blank identifier because we dont want to use index here
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
		}
	}               
}

//update handler
/*
	> set json Content-Type
	> params
	> loop over the movie , range
	> delete the movie with the ID that we have sent
	> add a new movie -the move that we send int the body

*/
func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application")
	params := mux.Vars(r)
	for index , item := range movies{
		if item.ID == params["id"]{
			movies =append(movies[:index], movies[index+1:]... )
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}


func main() {
	
	//initializing gorilla mux router
	r := mux.NewRouter()

	movies = append(movies,Movie{ID:"1", Isbn: "23112",Title: "Marvel 1",Director:&Director{FirstName:"Tony" , LastName: "Stark"}} )
	movies = append(movies,Movie{ID:"2", Isbn: "76546",Title: "Marvel 2",Director:&Director{FirstName:"Jhon" , LastName: "Smith"}} )
	movies = append(movies,Movie{ID:"3", Isbn: "43523",Title: "Marvel 3",Director:&Director{FirstName:"Dr." , LastName: "Strange"}} )

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}" , getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")


	fmt.Println("start server on port: 8000")
	err:=http.ListenAndServe(":8000",r)
	if err != nil {
		log.Fatal(err)
	}
}

