package spotify

import "fmt"

type MyTunes struct {
	playlist     PlayList
	nextPlaylist *MyTunes
	listHead     *MyTunes
}

func (t *MyTunes) AddNewPlaylist(playlist *PlayList) {
	var newNode *MyTunes = getNewNode()
	newNode.playlist = *playlist

	if getLastNode() == t.listHead {
		t.listHead = newNode
	} else {
		lastNode := getLastNode()
		lastNode.nextPlaylist = newNode
	}
}

func (t *MyTunes) ViewPlaylists() {
	for node := t.listHead; node != nil; node = node.nextPlaylist {
		fmt.Println("PlayList name : %s", node.playlist.playListName)

		func() {
			for _, value := range node.playlist.list {
				fmt.Printf("Song info - Artist : %s, Tile : %s, Duration : %d ", value.GetArtist(), value.GetDuration(), value.GetDuration())
			}
		}()
	}
}

// private
func getLastNode() (lastNode *MyTunes) {
	justAccess := MyTunes{}
	for head := justAccess.listHead; head != nil; head = head.nextPlaylist {
		lastNode = head
	}
	return
}

// private
func getNewNode() *MyTunes {
	return &MyTunes{}
}
