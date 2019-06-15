package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
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
		SystemHealth{ID: "1234", CPU: 23.65, Memory: 5.25},
		SystemHealth{ID: "3456", CPU: 25.65, Memory: 5.75},
	}
	vars := mux.Vars(r)
	key := vars["id"]
	if key == "" {
		// return all
		json.NewEncoder(w).Encode(systemHealthDataPoints)
	} else {
		for _, h := range systemHealthDataPoints {
			if h.ID == key {
				json.NewEncoder(w).Encode(h)
			}
		}
	}

	// fmt.Println(systemHealthDataPoints)
	// json.NewEncoder(w).Encode(systemHealthDataPoints)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am Home !")
}
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong !")
}

func handleRequest() {
	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/ping", ping)
	// http.HandleFunc("/systemhealth", systemHealthData)
	// fmt.Println("Listening to port 4050 ....")
	// log.Fatal(http.ListenAndServe(":4050", nil))
	// creates a new instance of a mux router
	//target := "https://httpbin.org"
	//remote, err := url.Parse(target)
	// if err != nil {
	// 	panic(err)
	// }

	//proxy := httputil.NewSingleHostReverseProxy(remote)

	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/ping", ping)
	myRouter.HandleFunc("/systemhealth/{id}", systemHealthData)
	myRouter.HandleFunc("/systemhealth", systemHealthData)
	myRouter.HandleFunc("/forward/url={targeturl}", handler)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":4050", myRouter))

}

func handler(w http.ResponseWriter, r *http.Request) {
	target := "http://www.google.com"
	vars := mux.Vars(r)
	key := vars["targeturl"]

	if key != "" {
		// return all
		target = "http://" + key
	}

	url, _ := url.Parse(target)
	fmt.Println(url)
	fmt.Println(url.Scheme)
	fmt.Println(url.Host)
	proxy := httputil.NewSingleHostReverseProxy(url)

	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = url.Host
	proxy.ServeHTTP(w, r)
}

func main() {

	fmt.Println("Rest API v2.0 - Mux Routers - listening in port 4050")
	handleRequest()
}
