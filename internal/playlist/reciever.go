package playlist

import (
	"fmt"
	"time"
)

func(p *Playlist) Reciever() {
	for {
		select {
		case data:= <-p.Nextprev:
			p.Stop <- struct{}{}
			if data == true {
				go p.Next()
			} else {
				go p.Prev()
			}
		case <-p.PlaySong:
			go p.Play()
		case data := <- p.Add:
			go p.AddSong(data)
		default:
			time.Sleep(time.Second)
			fmt.Printf("%s is playing\n", p.Current.Value.(Song).Title)
		}
	}
}