package spotify

import (
	"math/rand"
	"time"
)

type PlayList struct {
	playListName string
	list         map[int]Song
}

func (p *PlayList) AddSong(song *Song) {
	rand.Seed(time.Now().Unix())
	p.list[rand.Intn(100)] = *song
}

func (p *PlayList) TotalDuration() (duration int) {
	for _, value := range p.list {
		duration += value.GetDuration()
	}
	return
}

func (p *PlayList) CreateNewPlayList(playlistName string) {
	p.playListName = playlistName
	p.list = make(map[int]Song)
}
