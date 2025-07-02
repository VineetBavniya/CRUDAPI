package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, items := range movies {
		if items.ID == params["id"] {
			json.NewEncoder(w).Encode(items)
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movies)
}


func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, items := range movies {
		if items.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)

			var addMovie Movie
			_= json.NewDecoder(r.Body).Decode(&addMovie)
			addMovie.ID = params["id"]

			movies = append(movies, addMovie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}


func deleteMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, itmes := range movies {
		if itmes.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func main()  {
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "43532", Title: "Hera Pheri", Director : &Director{FirstName: "BaBu", LastName: "Rou"}})
	movies = append(movies, Movie{ID: "2", Isbn: "43532", Title: "Jon", Director: &Director{FirstName: "Jon", LastName: "Don"}})



	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	router.HandleFunc("/movies", createMovie).Methods("POST")

	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server At Port 8000\n")
	
	log.Fatal(http.ListenAndServe(":8000",router))
}