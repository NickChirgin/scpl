package handlers

import (
	"net/http"
)

func(h *Handler) Next(w http.ResponseWriter, r *http.Request){
	h.Pl.Nextprev <- true
	w.WriteHeader(http.StatusOK)
}