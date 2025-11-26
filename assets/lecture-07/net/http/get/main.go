package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// START OMIT

func main() {
	resp, err := http.Get("https://course-go.dev")
	if err != nil {
		log.Fatalf("could not execute request: %v", err)
	}

	fmt.Println(resp.Request.Proto)
	fmt.Println(resp.Request.Host)
	fmt.Println(resp.Request.URL)
	fmt.Println(resp.Request.Method)
	fmt.Println(resp.Request.Body)

	fmt.Println()

	fmt.Println(resp.Proto)
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	fmt.Println(resp.ContentLength)
	bodyBytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(bodyBytes))
}

// END OMIT
