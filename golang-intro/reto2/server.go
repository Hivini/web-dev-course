package main

import (
	"encoding/json"
	"net/http"
)

// Person refers to a living human
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func reto2(w http.ResponseWriter, r *http.Request) {
	persons := []Person{
		{
			Name:    "Vini",
			Age:     2,
			Address: "Sinaloa",
		},
		{
			Name:    "Ramon",
			Age:     31,
			Address: "Moshis",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func main() {
	fs := http.FileServer(http.Dir("./html"))

	http.HandleFunc("/reto2", reto2)
	http.Handle("/", fs)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
