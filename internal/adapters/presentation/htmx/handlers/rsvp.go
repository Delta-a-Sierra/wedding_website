package handlers

import (
	"net/http"

	"github.com/Delta-a-Sierra/wedding_website/internal/domain/app"
	"github.com/Delta-a-Sierra/wedding_website/internal/domain/entities"
)

type RSVPHandler struct {
	app *app.App
}

func NewRSVPHandler(app *app.App) *RSVPHandler {
	return &RSVPHandler{
		app: app,
	}
}

func (h *RSVPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	if name == "" || email == "" {
		w.Write([]byte("invalid data provided"))
		return
	}
	guest := entities.Guest{
		Email: email,
		Name:  name,
	}

	err := h.app.RSVP(r.Context(), guest)
	if err != nil {
		w.Write([]byte("internal error creating gues"))
		return
	}
}
