package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// START OMIT

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/articles/{date}-{slug}", getArticle)
	http.ListenAndServe(":3000", r)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	date := chi.URLParam(r, "date")
	slug := chi.URLParam(r, "slug")
	article, err := database.GetArticle(date, slug)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(fmt.Sprintf("error fetching article %s-%s: %v", date, slug, err)))
		return
	}

	w.Write([]byte(article.Text()))
}

// END OMIT
