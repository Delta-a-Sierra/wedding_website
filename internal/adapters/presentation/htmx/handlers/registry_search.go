package handlers

import (
	"net/http"

	"github.com/Delta-a-Sierra/wedding_website/internal/adapters/presentation/htmx/templates/sections"
	"github.com/Delta-a-Sierra/wedding_website/internal/domain/app"
	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
)

type GetRegistryHandler struct {
	app *app.App
}

func NewGetRegistryHandler(app *app.App) *GetRegistryHandler {
	return &GetRegistryHandler{
		app: app,
	}
}

func (h *GetRegistryHandler) SearchAll(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	searchString := r.Form.Get("registry-search")
	var items []entities.RegistryItem
	var err error
	if searchString == "" {
		items, err = h.app.GetRegistryItems(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		sections.RegistryItemGrid(items).Render(r.Context(), w)
		return
	}
	items, err = h.app.SearchRegistry(r.Context(), searchString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	sections.RegistryItemGrid(items).Render(r.Context(), w)
}

func (h *GetRegistryHandler) SearchNotPurchased(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	searchString := r.Form.Get("registry-search")
	var items []entities.RegistryItem
	var err error
	if searchString == "" {
		items, err = h.app.GetRegistryItemsNotPurchased(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		sections.RegistryItemGrid(items).Render(r.Context(), w)
		return
	}
	items, err = h.app.SearchRegistryNotPurchased(r.Context(), searchString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	sections.RegistryItemGrid(items).Render(r.Context(), w)
}

func (h *GetRegistryHandler) FilterAll(w http.ResponseWriter, r *http.Request) {
	items, err := h.app.GetRegistryItems(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	sections.RegistryItems(items, true).Render(r.Context(), w)
}

func (h *GetRegistryHandler) FilterNotPurchased(w http.ResponseWriter, r *http.Request) {
	items, err := h.app.GetRegistryItemsNotPurchased(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	sections.RegistryItems(items, false).Render(r.Context(), w)
}
