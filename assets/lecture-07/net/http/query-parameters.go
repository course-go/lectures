package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/users", handleUser)

	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(w, "Hello %s\n", name)
	} else {
		fmt.Fprintf(w, "Mr. Nobody\n")
	}
}

// END OMIT
