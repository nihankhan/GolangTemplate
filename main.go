package main

import (
	"log"
	"net/http"

	"github.com/nihankhan/GolangTemplate/handlers"
)

const (
	PORT  = ":8080"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Query)


	log.Println("Server listening on http://127.0.0.1" + PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))	
}

