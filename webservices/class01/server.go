package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
)

type Song struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Views  int    `json:"views"`
}

var Songs []Song

func returnAllSongs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllSongs")
	json.NewEncoder(w).Encode(Songs)
}

func returnTopFiveSongs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnTopFiveSongs")
	sorted, err := strconv.ParseBool(r.Header.Get("Sorted"))
	if err != nil {
		log.Fatalln("Problems...")
	}
	limit, err := strconv.Atoi(r.Header.Get("Limit"))
	if err != nil {
		limit = -1
	}
	copyOfSongs := make([]Song, len(Songs))
	copy(copyOfSongs, Songs)
	if sorted {
		sort.Slice(copyOfSongs, func(i, j int) bool {
			return copyOfSongs[i].Views > copyOfSongs[j].Views
		})
	}

	limited := make([]Song, limit)
	for i := 0; i < limit && i < len(copyOfSongs); i++ {
		limited[i] = copyOfSongs[i]
	}
	json.NewEncoder(w).Encode(limited)
}

func main() {
	Songs = []Song{
		{Title: "Song 1", Author: "Author 1", Views: 50},
		{Title: "Song 2", Author: "Author 2", Views: 15},
		{Title: "Song 3", Author: "Author 3", Views: 6},
		{Title: "Song 4", Author: "Author 4", Views: 100},
		{Title: "Song 5", Author: "Author 5", Views: 45},
		{Title: "Song 6", Author: "Author 6", Views: 30},
		{Title: "Song 7", Author: "Author 7", Views: 5},
		{Title: "Song 8", Author: "Author 8", Views: 20},
	}
	fs := http.FileServer(http.Dir("./html"))

	http.Handle("/", fs)
	http.HandleFunc("/songs", returnAllSongs)
	http.HandleFunc("/songs/top5", returnTopFiveSongs)

	log.Println("Starting server...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println(err)
	}
}
