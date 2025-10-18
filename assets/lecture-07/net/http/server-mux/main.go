package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Homepage\n\n")
		fmt.Fprintf(w, "You are visiting from: %s", r.RemoteAddr)
	})
	router.HandleFunc("/status", handleStatus)

	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I'm fine!")
}

// END OMIT
