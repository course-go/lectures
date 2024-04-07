package main

import (
	"fmt"
	"net/http"
	"time"
)

// START OMIT

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Homepage\n\n")
		fmt.Fprintf(w, "You are visiting from: %s", r.RemoteAddr)
	})
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}
	server.ListenAndServe()
}

// END OMIT
