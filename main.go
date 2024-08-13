package main

import (
	"log"
	"net/http"

	"functions/functions"
)

func main() {
	http.HandleFunc("/static/", functions.SetupStaticFilesHandlers)
	http.HandleFunc("/", functions.Handler)
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
