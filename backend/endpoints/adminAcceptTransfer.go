package endpoints

import (
	"WDB/views"
	"net/http"
)

func (h *Handlers) AdminAcceptTransfer(w http.ResponseWriter, r *http.Request) {
	views.ResponseWithJSON(w, 200, h.transactions)
}
