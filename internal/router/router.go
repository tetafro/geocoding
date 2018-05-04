package router

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/tetafro/geocoding/internal/places"
)

// New creates main HTTP handler for the app.
func New(p *places.Controller) http.Handler {
	r := chi.NewRouter()
	r.Get("/", home)
	r.Get("/healthz", healthz)
	r.Get("/api/v1/places", p.Get)

	// Static
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	r.Get("/static/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))

	return r
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
