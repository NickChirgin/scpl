package main

import (
	"fmt"

	"github.com/nickchirgin/scpl/internal/playlist"
)

func main() {
	pl := playlist.NewPlaylist()
	go pl.AddSong(playlist.Song{Title: "Yo", Duration: 10})
	select {
	case data:=<-pl.PlaySong:
		fmt.Println(data)
	}
}