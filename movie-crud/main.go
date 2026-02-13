package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// "encoding/json"

// "math/rand"
// "net/http"
// "strconv"

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {

			json.NewEncoder(w).Encode(item)
			return

		}
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "movie not found",
	})

}

func createMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for idx, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "movie not found",
	})

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for idx, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)

}

var movies []Movie

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:       "1",
		Isbn:     "438227",
		Title:    "Movie One",
		Director: &Director{Firstname: "John", Lastname: "Doe"},
	})

	movies = append(movies, Movie{
		ID:       "2",
		Isbn:     "454555",
		Title:    "Movie Two",
		Director: &Director{Firstname: "Steve", Lastname: "Smith"},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("The go server is running ")

	log.Fatal(http.ListenAndServe(":8000", r))

}
