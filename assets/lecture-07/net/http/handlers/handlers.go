package handlers

import "net/http"

// START OMIT

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// END OMIT
