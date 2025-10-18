package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome!")
	})

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":8080", n)
}
