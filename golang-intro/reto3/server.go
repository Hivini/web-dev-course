package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// SafeData is safe.
type SafeData struct {
	mu   sync.Mutex
	data []int
}

// Data is the main structure, pretty simple
type Data struct {
	Arr []int `json:"arr"`
}

var sd = SafeData{
	data: make([]int, 0),
}

// GetData gets the value of the field "data"
func (d *SafeData) GetData() []int {
	sd.mu.Lock()
	// Unlock after finishing return
	defer sd.mu.Unlock()
	return sd.data
}

// AddData adds a value to the field "data" and returns the new value
func (d *SafeData) AddData(num int) []int {
	sd.mu.Lock()
	// Unlock after finishing return
	defer sd.mu.Unlock()
	sd.data = append(sd.data, num)
	return sd.data
}

func reto3(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	counter := r.Form.Get("counter")
	log.Printf("User counter: %s", counter)
	value, err := strconv.Atoi(counter)
	if err != nil {
		fmt.Println("Algo murio")
	}

	// Depending on the counter given by the user, we will give the rest of the slice
	newData := Data{
		Arr: sd.GetData()[value:],
	}
	log.Printf("New data sent to the user: %v", newData.Arr)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newData)
}

func dataProvider() {
	for {
		time.Sleep(3 * time.Second)
		newData := sd.AddData(rand.Intn(100))
		log.Printf("New data value: %v", newData)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./html"))

	http.HandleFunc("/reto3", reto3)
	http.Handle("/", fs)

	go dataProvider()

	log.Println("Starting server...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
