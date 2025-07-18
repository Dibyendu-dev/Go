package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("welcome to mod in golang")
	greet()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000",r))
}

func greet() {
	fmt.Println("Hello dibyendu")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> welcome to golang </h1>"))
}
