package handlers

import (
	"fmt"
	"net/http"

	"github.com/nickchirgin/scpl/internal/playlist"
)

func(h *Handler) Pause(w http.ResponseWriter, r *http.Request){
	h.Pl.Stop <- struct{}{}
	fmt.Printf("Song %s stopped at %d", h.Pl.Current.Value.(playlist.Song).Title, h.Pl.Time)
	w.WriteHeader(http.StatusOK)
}