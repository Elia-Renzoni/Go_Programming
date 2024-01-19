package spotify

type Song struct {
	title, artist string
	duration      int
}

func (s *Song) CreateNewSong(title, artist string, duration int) {
	s.title = title
	s.artist = artist
	s.duration = duration
}

func (s *Song) GetTile() string {
	return s.title
}

func (s *Song) GetArtist() string {
	return s.artist
}

func (s *Song) GetDuration() int {
	return s.duration
}
