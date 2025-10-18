package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/users/{id}", handleUser)
	router.HandleFunc("/users/bob", handleBob)

	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "You just looked up user with ID: %s", id)
}

func handleBob(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You just looked up Bob!")
}

// END OMIT
