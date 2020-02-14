package rest

import "net/http"

func (h *Handler) GetPing(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, 200, "Pong")
}
