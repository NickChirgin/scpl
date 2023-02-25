package handlers

import (
	"net/http"

	"github.com/nickchirgin/scpl/internal/playlist"
)
type Handler struct {
	Pl *playlist.Playlist
}
func (h *Handler)  Play(w http.ResponseWriter, r *http.Request) {
	h.Pl.PlaySong <- struct{}{}
	w.WriteHeader(http.StatusOK)
}
