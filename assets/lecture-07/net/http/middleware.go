package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/users/{id}", handleUser)

	// Middleware
	handler := logging(router)

	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "You just looked up user with ID: %s", id)
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL, time.Now())
		next.ServeHTTP(w, r)
	})
}
