package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nickchirgin/scpl/internal/playlist"
)
type Handler struct {
	Pl *playlist.Playlist
}
func (h *Handler)  Play(w http.ResponseWriter, r *http.Request) {
	h.Pl.PlaySong <- struct{}{}
	fmt.Printf("%s is playing from %d", h.Pl.Current.Value.(playlist.Song).Title, h.Pl.Time)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Base(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	s, _ := json.Marshal("rabotaet")
	w.Write(s)
}