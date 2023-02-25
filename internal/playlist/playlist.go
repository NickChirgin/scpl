package playlist

import (
	"container/list"
	"context"
	"fmt"
	"sync"
	"time"
)

type Playlist struct {
	 Tracks *list.List
	 Current *list.Element
	 Time int
	 mu sync.RWMutex
	 Add chan Song 
	 Nextprev chan bool  // true = next, false = prev
	 PlaySong chan struct{} 
	 Stop chan struct{}
}

type Song struct {
	Title string `json:"title"` 
	Duration int `json:"duration"`
}

func NewPlaylist() Playlist {
	return Playlist{Tracks: list.New(), Current: &list.Element{}, PlaySong: make(chan struct{}), Stop: make(chan struct{}), Nextprev: make(chan bool), Time: 0, Add: make(chan Song)}
}

func (p *Playlist) Play() {
	p.mu.RLock()
	defer p.mu.RUnlock() 
	if p.Tracks.Len() == 0 {
		fmt.Println("Playlist is empty")
		return
	}
	if p.Current == nil {
		p.Current = p.Tracks.Front()
	}
	if p.Time ==  0{
		p.Time = p.Current.Value.(Song).Duration
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(p.Time) * time.Second)
	defer cancel() 
	deadline, _ := ctx.Deadline()
	fmt.Printf("%s is now playing \n", p.Current.Value.(Song).Title)
	select {
	case <-ctx.Done():
		fmt.Printf("%s song has ended\n", p.Current.Value.(Song).Title)
		p.Next()
	case <-p.Stop: 
		p.Time = int(time.Until(deadline).Seconds())
		break
	}	
}

func (p *Playlist) Pause() {
	p.Stop <- struct{}{}
}

func (p *Playlist) AddSong(track Song) {
	p.mu.Lock()
	defer p.mu.Unlock()
	fmt.Printf("%s song has been added\n", track.Title)
	p.Tracks.PushBack(track)
}

func (p *Playlist) Next() {
	p.mu.RLock()
	defer p.mu.RUnlock()
	p.Time = 0
	if p.Current.Next() != nil || p.Current.Next().Value != (Song{}) {
		p.Current = p.Current.Next()
		p.Play()
	}
}

func (p *Playlist) Prev() {
	p.mu.RLock()
	defer p.mu.RUnlock()
	p.Time = 0
	if p.Current.Prev() != nil {
		p.Current = p.Current.Prev()
		p.Play()
	}
}