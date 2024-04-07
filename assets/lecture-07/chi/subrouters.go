package main

import "github.com/go-chi/chi/v5"

// START OMIT

func main() {
	r := chi.NewRouter()
	r.Route("/articles", func(r chi.Router) {
		r.With(paginate).Get("/", listArticles)                           // GET /articles
		r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017
		r.Post("/", createArticle)                                        // POST /articles
		r.Get("/search", searchArticles)                                  // GET /articles/search

		// Regexp url parameters:
		r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug) // GET /articles/home-is-toronto

		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(ArticleCtx)
			r.Get("/", getArticle)       // GET /articles/123
			r.Put("/", updateArticle)    // PUT /articles/123
			r.Delete("/", deleteArticle) // DELETE /articles/123
		})
	})
}

// END OMIT
