package handlers

import (
	"fmt"
	"net/http"

	"github.com/Delta-a-Sierra/wedding_website/internal/adapters/presentation/htmx/templates/components"
	"github.com/Delta-a-Sierra/wedding_website/internal/domain/app"
)

type GetCountdownHandler struct {
	app *app.App
}

func NewGetCountdownHandler(app *app.App) *GetCountdownHandler {
	return &GetCountdownHandler{
		app: app,
	}
}

func (h *GetCountdownHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	countdownm, err := h.app.GetCountdown(r.Context())
	countdownm.ZeroOutMinusValues()
	if err != nil {
		fmt.Fprintf(w, "err: %s", err)
		return
	}
	components.Countdown(countdownm).Render(r.Context(), w)
}
