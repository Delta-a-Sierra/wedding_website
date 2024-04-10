package handlers

import (
	"net/http"

	"github.com/Delta-a-Sierra/wedding_website/internal/adapters/presentation/htmx/templates/sections"
	"github.com/Delta-a-Sierra/wedding_website/internal/domain/app"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type GetRegistryItemPurchasedHandler struct {
	app *app.App
}

func NewGetRegistryItemPurchasedHandler(app *app.App) *GetRegistryItemPurchasedHandler {
	return &GetRegistryItemPurchasedHandler{
		app: app,
	}
}

func (h *GetRegistryItemPurchasedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uid, _ := uuid.Parse(id)
	h.app.DelcareRegistryPurchase(r.Context(), uid, uuid.New())
	item, err := h.app.GetRegistryItem(r.Context(), uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	item.Purchased = true
	sections.RegistryItem(item).Render(r.Context(), w)
}
