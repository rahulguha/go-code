package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am Home !")
}
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong !")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe(":4050", nil))
}

func main() {
	handleRequest()
}
