package handlers

import "net/http"

func(h *Handler) Prev(w http.ResponseWriter, r *http.Request){
	h.Pl.Nextprev <- false
	w.WriteHeader(http.StatusOK)
}