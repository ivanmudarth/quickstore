package main

import (
	"fmt"
	"net/http"

	"./database"
	"./handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func init() {
	handlers.CreateAWSConfig()
	database.DBInit()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/upload", handlers.UploadHandler).Methods("POST")
	r.HandleFunc("/display", handlers.DisplayHandler).Methods("GET")
	http.Handle("/", r)

	// Enable CORS for client
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	fmt.Println("Starting server at port 8080...\n")
	http.ListenAndServe(":8080", handler)
}
