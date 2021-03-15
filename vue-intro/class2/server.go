package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Pelicula refers to a movie.
type Pelicula struct {
	Name     string `json:"name"`
	Director string `json:"director"`
}

func retote(w http.ResponseWriter, r *http.Request) {
	peliculas := []Pelicula{
		{
			Name:     "El Chapo Mejorado",
			Director: "Mark Anthony",
		},
		{
			Name:     "Terminator",
			Director: "Michael Bay",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peliculas)
}

func main() {
	fs := http.FileServer(http.Dir("./html"))

	http.Handle("/", fs)
	http.HandleFunc("/getJson", retote)

	log.Println("Starting server...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
