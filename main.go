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
	ID          string    `json:"id"`
	Isbn        string    `json:"isbn"`
	Title       string    `json:"title"`
	Director    *Director `json:"director"`
	Genres      []string  `json:"genres"`
	Rating      float64   `json:"rating"`
	Streaming   string    `json:"streaming"`
	Description string    `json:"description"`
	Cast        []string  `json:"cast"`
}

type Director struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Nationality string `json:"nationality"`
	Age         int    `json:"birthyear"`
	NetWorth    string `json:"networth"`
}

var movies []Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	// TODO:{
	// set json content type
	//params
	//loop over the movies, range
	//delete the movie with id that you've sent
	// add a new movie that we send in the body of the postmon

	//}
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	r := mux.NewRouter()
	// Add your routes here
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "438222",
		Title: "Iron Man",
		Director: &Director{
			FirstName:   "Sahil",
			LastName:    "Dhargave",
			Nationality: "Indian",
			Age:         34,
			NetWorth:    "$234",
		},
		Genres:      []string{"Action", "Sci-Fi", "Mystery"},
		Rating:      4.5,
		Streaming:   "Netflix",
		Description: "Tony Stark. Genius, billionaire, playboy, philanthropist. Son of legendary inventor and weapons contractor Howard Stark. When Tony Stark is assigned to give a weapons presentation to an Iraqi unit led by Lt. Col. James Rhodes, he's given a ride on enemy lines. That ride ends badly when Stark's Humvee that he's riding in is attacked by enemy combatants. He survives - barely - with a chest full of shrapnel and a car battery attached to his heart. In order to survive he comes up with a way to miniaturize the battery and figures out that the battery can power something else. Thus Iron Man is born. He uses the primitive device to escape from the cave in Iraq. Once back home, he then begins work on perfecting the Iron Man suit. But the man who was put in charge of Stark Industries has plans of his own to take over Tony's technology for other matters.—halo1k",
		Cast: []string{
			"Robert Downey Jr.",
			"Terrence Howard",
			"Gwyneth Paltrow",
			"Leslie Bibb",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "438225",
		Title: "Traffic",
		Director: &Director{
			FirstName:   "Ram",
			LastName:    "Trivadi",
			Nationality: "Indian",
			Age:         54,
			NetWorth:    "$453",
		},
		Genres:      []string{"Adventure", "Triller", "Drama"},
		Rating:      4.5,
		Streaming:   "Amazon",
		Description: "An emotional thriller based on a road trip from Mumbai to Pune. Inspired from the real events that happened in Chennai.",
		Cast: []string{
			"Manoj Bajpayee",
			"Jimmy Shergill",
			"Divya Dutta",
		},
	})

	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovieByID).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
