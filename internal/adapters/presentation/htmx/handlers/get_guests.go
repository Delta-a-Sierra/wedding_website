package handlers

import (
	"fmt"
	"net/http"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/app"
)

type GetGuestsHandler struct {
	app *app.App
}

func NewGetGuestsHandler(app *app.App) *GetGuestsHandler {
	return &GetGuestsHandler{
		app: app,
	}
}

func (h *GetGuestsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	guests, _ := h.app.GetGuestList(r.Context())
	w.Write([]byte(fmt.Sprintf("%+v", guests)))
}
