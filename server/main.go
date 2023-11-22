package main

import (
	"fmt"
	"net/http"

	"./handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func init() {
	handlers.CreateAWSConfig()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/upload", handlers.UploadHandler).Methods("POST")
	r.HandleFunc("/display", handlers.DisplayHandler).Methods("GET")
	http.Handle("/", r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	fmt.Println("Starting server at port 8080...\n")
	http.ListenAndServe(":8080", handler)
}

// there must be only one package per folder
// variables and functions can be exported by capitalizing the first letter
// camelCase used for variable naming

/*
upload file
- store metadata about file in postgres
	- id, user (session ?), filename, user tags, auto tags, size
- store file in s3

retrieve file

*/
