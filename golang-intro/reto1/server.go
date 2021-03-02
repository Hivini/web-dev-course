package main

import (
	"fmt"
	"net/http"
)

func darMensaje(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s\n", r.Cookies())
}

func main() {
	http.HandleFunc("/", darMensaje)
	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}
