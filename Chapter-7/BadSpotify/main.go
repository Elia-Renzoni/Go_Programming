/**
*	@author Elia Renzoni
*	@date 19/01/2024
*	@brief Naive Spotify
*	TO END
**/

package main

import (
	"fmt"
	"spotiymodule/spotify"
)

func main() {
	var (
		songTitle, songAuthor, playListName string
		songDuration                        int
	)

	songTitle, songAuthor, songDuration = getSongInfo()
	song1 := &spotify.Song{}
	song1.CreateNewSong(songTitle, songAuthor, songDuration)
	songTitle, songAuthor, songDuration = getSongInfo()
	song2 := &spotify.Song{}
	song2.CreateNewSong(songTitle, songAuthor, songDuration)

	playListName = getPlayListName()
	play1 := &spotify.PlayList{}
	play1.CreateNewPlayList(playListName)

	play1.AddSong(song1)
	play1.AddSong(song2)

	fmt.Printf("Playlist total duration : %d", play1.TotalDuration())

	myTune := &spotify.MyTunes{}
	myTune.AddNewPlaylist(play1)
	myTune.ViewPlaylists()
}

func getSongInfo() (title string, author string, duration int) {
	var (
		control bool = true
		err     error
	)
	for control {
		fmt.Printf("Inerisci le info : ")
		_, err = fmt.Scanf("%s", &title)
		_, err = fmt.Scanf("%s", &author)
		_, err = fmt.Scanf("%d", &duration)

		if err != nil {
			return "", "", 0
		} else {
			control = false
		}
	}
	return
}

func getPlayListName() (name string) {
	var err error
	var control bool = true
	for control {
		fmt.Printf("Ïnserisci il nome della playlist")
		_, err = fmt.Scanf("%s", &name)
		if err != nil {
			return ""
		} else {
			control = false
		}
	}
	return
}
