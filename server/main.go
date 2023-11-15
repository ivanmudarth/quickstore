package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	http.Handle("/", r)
	r.HandleFunc("/hello", handler).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// there must be only one package per folder
// variables and functions can be exported by capitalizing the first letter
// camelCase used for variable naming

/*
upload file

retrieve file

*/
