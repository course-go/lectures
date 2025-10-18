package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// START OMIT

type PostRequest struct {
	Data string `json:"data"`
}

func main() {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	payload := PostRequest{Data: "Hello there!"}
	payloadJson, _ := json.Marshal(payload)
	fmt.Println(string(payloadJson))

	reader := bytes.NewReader([]byte(payloadJson))
	req, _ := http.NewRequest(http.MethodPost, "http://course-go.dev/upload", reader)
	req.Header.Add("Content-Type", "application/json")
	fmt.Println(req)

	resp, _ := client.Do(req)

	fmt.Println(resp.Status)
}

// END OMIT
