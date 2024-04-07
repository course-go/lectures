package main

import (
	"log"
	"net/http"
)

// START OMIT

func main() {
	router := http.NewServeMux()
	router.HandleFunc("POST /users/{id}", handlePostUser)
	router.HandleFunc("DELETE /users/{id}", handleDeleteUser)
	router.HandleFunc("/users/{id}", handleRest) // Handles rest of the methods
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handlePostUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func handleRest(w http.ResponseWriter, r *http.Request) {
	switch r.Method { // Handling before Go 1.22
	case http.MethodGet:
	case http.MethodPut:
	}
}

// END OMIT
