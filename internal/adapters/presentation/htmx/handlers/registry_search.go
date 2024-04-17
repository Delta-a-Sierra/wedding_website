package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Delta-a-Sierra/wedding_website/internal/adapters/presentation/htmx/templates/sections"
	"github.com/Delta-a-Sierra/wedding_website/internal/domain/app"
	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
	"github.com/go-chi/chi"
)

type GetRegistryHandler struct {
	app *app.App
}

func NewGetRegistryHandler(app *app.App) *GetRegistryHandler {
	return &GetRegistryHandler{
		app: app,
	}
}

func (h *GetRegistryHandler) GetRegistryPageFilteredAll(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	fmt.Println("page", page)
	i, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if i == 0 {
		http.Error(w, "page is 0", http.StatusInternalServerError)
		return
	}
	pages, err := h.app.GetPageCount(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	items, err := h.app.GetRegistryItemsPageAll(r.Context(), 6, i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("page", i)
	if err = sections.RegistryItems(items, true, pages, i).Render(r.Context(), w); err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *GetRegistryHandler) GetRegistryPageFilteredNotPurchased(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	i, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if i == 0 {
		http.Error(w, "page is 0", http.StatusInternalServerError)
		return
	}

	pages, err := h.app.GetPageCountNotPurchased(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	items, err := h.app.GetRegistryItemsPageNotPurchased(r.Context(), 6, i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = sections.RegistryItems(items, false, pages, i).Render(r.Context(), w); err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *GetRegistryHandler) SearchAll(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	searchString := r.Form.Get("registry-search")
	var items []entities.RegistryItem
	var err error
	if searchString == "" {
		pages, err := h.app.GetPageCount(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		items, err = h.app.GetRegistryItemsPageAll(r.Context(), 6, 1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Println(items)

		sections.RegistryItemGrid(items, pages, 1, "/registry/all/page").Render(r.Context(), w)
		return
	}
	items, err = h.app.SearchRegistry(r.Context(), searchString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	pages, err := h.app.GetPageCount(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sections.RegistryItemGrid(items, pages, 1, "/registry/all/page").Render(r.Context(), w)
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

		pages, err := h.app.GetPageCountNotPurchased(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sections.RegistryItemGrid(items, pages, 1, "/registry/not-purchased/page").Render(r.Context(), w)
		return
	}
	items, err = h.app.SearchRegistryNotPurchased(r.Context(), searchString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	pages, err := h.app.GetPageCountNotPurchased(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sections.RegistryItemGrid(items, pages, 1, "/registry/not-purchased/page").Render(r.Context(), w)
}

func (h *GetRegistryHandler) FilterAll(w http.ResponseWriter, r *http.Request) {
	items, err := h.app.GetRegistryItems(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pages, err := h.app.GetPageCount(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sections.RegistryItems(items, true, pages, 1).Render(r.Context(), w)
}

func (h *GetRegistryHandler) FilterNotPurchased(w http.ResponseWriter, r *http.Request) {
	items, err := h.app.GetRegistryItemsNotPurchased(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pages, err := h.app.GetPageCountNotPurchased(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sections.RegistryItems(items, false, pages, 1).Render(r.Context(), w)
}
