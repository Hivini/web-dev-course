package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseOne struct {
	TestArr [5]int `json:"testArr"`
}

func main() {
	http.HandleFunc("/coolarray", ServeArray)
	http.HandleFunc("/", MainHtml)
	log.Printf("Starting server on localhost:8080...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MainHtml(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func ServeArray(w http.ResponseWriter, r *http.Request) {
	log.Printf("request for array: %s", r.Host)
	arr := [5]int{1, 2, 3, 4, 5}
	w.Header().Set("Content-Type", "application/json")
	response := ResponseOne{
		TestArr: arr,
	}
	json.NewEncoder(w).Encode(response)
}
