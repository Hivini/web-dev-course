package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/eljson", ServeNoJson)
	http.HandleFunc("/secondfile", ServeSecondFile)
	http.HandleFunc("/testjson", ServeTestJson)
	http.HandleFunc("/", MainHtml)
	log.Printf("Starting server on localhost:8080...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MainHtml(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func ServeNoJson(w http.ResponseWriter, r *http.Request) {
	log.Printf("request nojson file: %s", r.Host)
	http.ServeFile(w, r, "nojson.txt")
}

func ServeSecondFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("request secondfile: %s", r.Host)
	http.ServeFile(w, r, "secondfile.txt")
}

func ServeTestJson(w http.ResponseWriter, r *http.Request) {
	log.Printf("request testjson: %s", r.Host)
	http.ServeFile(w, r, "testjson.json")
}
