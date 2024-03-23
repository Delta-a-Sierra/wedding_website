package htmx

import (
	"net/http"

	"github.com/Delta-a-Sierra/wedding_website/internal/adapters/presentation/htmx/templates"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Home("test").Render(r.Context(), w)
	})
	http.ListenAndServe(":3000", r)
}
