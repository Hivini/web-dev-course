package main

import (
	//"fmt"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var reto1array [5]int

type book struct {
	Titulo string `json: titulo`
	Autor  string `json: autor`
}

func mostrarHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ajax02.html")
}

func darMensaje(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	y := r.Form.Get("y")
	x := r.Form.Get("x")

	fmt.Printf("x = %s\n", x)
	fmt.Printf("y = %s\n", y)

	libro := book{
		Titulo: "La Casa",
		Autor:  "Paco Roca",
	}

	// w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(libro)
}

func reto1(w http.ResponseWriter, r *http.Request) {
	requestType := r.Method
	switch requestType {
	case "GET":
		fmt.Fprint(w, reto1array)
	case "POST":
		r.ParseForm()
		x := r.Form.Get("x")
		pos := r.Form.Get("pos")
		value, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println("Algo murio")
			return
		}
		posValue, err := strconv.Atoi(pos)
		if err != nil {
			fmt.Println("Error en la posicion")
			return
		}
		// Suponemos que siempre la posicion esta bien
		reto1array[posValue] = value
		w.WriteHeader(http.StatusOK)
	default:
		fmt.Printf("WHAAAT")
	}
}

func main() {
	reto1array = [5]int{1, 2, 3, 4, 5}

	http.HandleFunc("/dato", darMensaje)
	http.HandleFunc("/reto1", reto1)
	http.HandleFunc("/", mostrarHTML)

	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}
