package htmx

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Delta-a-Sierra/wedding_website/internal/adapters/presentation/htmx/handlers"
	"github.com/Delta-a-Sierra/wedding_website/internal/adapters/presentation/htmx/templates/layout"
	"github.com/Delta-a-Sierra/wedding_website/internal/adapters/presentation/htmx/templates/pages"
	"github.com/Delta-a-Sierra/wedding_website/internal/domain/app"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Date struct {
	Days int
}

func Start(app *app.App) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	fileServer := http.FileServer(http.Dir("internal/adapters/presentation/htmx/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		countdown, err := app.GetCountdown(r.Context())
		if err != nil {
			fmt.Fprintf(w, "err: %s", err.Error())
			return
		}
		countdown.ZeroOutMinusValues()

		items, err := app.GetRegistryItems(r.Context())
		if err != nil {
			fmt.Fprintf(w, "err: %s", err.Error())
			return
		}
		layout.Base("wedding_website", pages.Home(countdown, items)).Render(r.Context(), w)
	})
	r.Get("/countdown", handlers.NewGetCountdownHandler(app).ServeHTTP)
	log.Fatal(http.ListenAndServe(":3000", r))
}
