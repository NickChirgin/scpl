package playlist

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
		}
	}
}