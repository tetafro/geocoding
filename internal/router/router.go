package router

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/tetafro/geocoding/internal/places"
)

// New creates main HTTP handler for the app.
func New(p *places.Controller) http.Handler {
	r := chi.NewRouter()
	r.Get("/healthz", Healthz)
	r.Get("/api/places", p.Get)
	return r
}

// Healthz simply gives a 200 reply.
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}
