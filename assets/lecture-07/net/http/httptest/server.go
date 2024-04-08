package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// START OMIT

func TestSignatureEditionBoxThreeClient(t *testing.T) {
	expected := "{'data':{'name':'Gavin Belson'}}"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer server.Close()

	client := NewHooliClient(server.URL) // Server hostname can easily be accessed
	response, err := client.ExecuteRequest()
	if err != nil {
		t.Fatalf("error should nil but got %v", err)
	}

	actual := strings.TrimSpace(response)
	if actual != expected {
		t.Errorf("response should be %s but got %s", expected, actual)
	}
}

// END OMIT
