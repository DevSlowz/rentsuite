package main

import (
	"log"
	"net/http"
)

// Home Handler which will write bytes containing the message
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Rent Suite"))
}

func main() {

	// Initialize a new servemux - Router
	mux := http.NewServeMux()
	// "/" URL pattern will be handled by the home handler
	mux.HandleFunc("/", home)

	// Print message indicating it is live
	log.Print("Starting server on: 4000")

	// If an error occurs capture and log it
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
