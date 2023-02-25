package handlers

import (
	"fmt"
	"net/http"

	"github.com/nickchirgin/scpl/internal/playlist"
)

func(h *Handler) Pause(w http.ResponseWriter, r *http.Request){
	h.Pl.Stop <- struct{}{}
	fmt.Printf("%s song has been stopped\n", h.Pl.Current.Value.(playlist.Song).Title)
	w.WriteHeader(http.StatusOK)
}