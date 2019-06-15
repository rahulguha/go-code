package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SystemHealth struct {
	ID     string  "json:'id'"
	CPU    float32 "json:'cpu'"
	Memory float32 "json:'memory'"
}
type SystemHealthDataPoints []SystemHealth

func systemHealthData(w http.ResponseWriter, r *http.Request) {
	// Mock SystemHealth for now
	systemHealthDataPoints := SystemHealthDataPoints{
		SystemHealth{ID: "A1234", CPU: 23.65, Memory: 5.25},
		SystemHealth{ID: "A1235", CPU: 25.65, Memory: 5.75},
	}
	json.NewEncoder(w).Encode(systemHealthDataPoints)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am Home !")
}
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong !")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/systemhealth", systemHealthData)
	log.Fatal(http.ListenAndServe(":4050", nil))
}

func main() {
	handleRequest()
}
