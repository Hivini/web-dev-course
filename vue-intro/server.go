package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./html"))

	http.Handle("/", fs)

	log.Println("Starting server...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
