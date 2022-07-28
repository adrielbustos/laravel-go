package handers

import (
	"net/http"

	"github.com/tsawler/celeritas"
)

type Handers struct {
	App *celeritas.Celeritas
}

func (h *Handers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering: ", err)
	}
}
