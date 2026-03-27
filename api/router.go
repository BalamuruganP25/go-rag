package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(handler *Handler) http.Handler {

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		// r.Get("/health", HealthHandler)
		r.Post("/ask", handler.AskHandler)
		r.Post("/documents", DocumentHandler)
		r.Post("/documents/pdf", UploadPDFHandler)

	})

	return r
}
