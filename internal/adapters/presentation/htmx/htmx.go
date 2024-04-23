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

		p, err := app.GetPageCount(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		items, err := app.GetRegistryItemsPageAll(r.Context(), 6, 1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		layout.Base("wedding_website", pages.Home(countdown, items, p, 1)).Render(r.Context(), w)
	})
	registryHandlers := handlers.NewGetRegistryHandler(app)
	r.Get("/countdown", handlers.NewGetCountdownHandler(app).ServeHTTP)
	r.Post("/rsvp", handlers.NewRSVPHandler(app).ServeHTTP)
	r.Post("/registry/search", registryHandlers.SearchAll)
	r.Post("/registry/search/not-purchased", registryHandlers.SearchNotPurchased)
	r.Get("/registry/all", registryHandlers.FilterAll)
	r.Get("/maps", registryHandlers.GetMapScript)
	r.Get("/registry/all/page/{page}", registryHandlers.GetRegistryPageFilteredAll)
	r.Get("/registry/not-purchased/page/{page}", registryHandlers.GetRegistryPageFilteredNotPurchased)
	r.Get("/registry/not-purchased", registryHandlers.FilterNotPurchased)
	r.Post("/registry/purchased/{id}", handlers.NewGetRegistryItemPurchasedHandler(app).ServeHTTP)
	r.Get("/admin/guests", handlers.NewGetGuestsHandler(app).ServeHTTP)
	log.Fatal(http.ListenAndServe(":3000", r))
}
