package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	const filepathRoot = "."
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	registerPaths(mux)
	server := NewServer(port, mux)
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
