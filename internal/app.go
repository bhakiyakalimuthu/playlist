package internal

import "playlist/internal/linkedlist"

func SinglyPlayList()  {

	pl := linkedlist.NewPlayListS("singly playlist")
	pl.Add("song1","artist1")
	pl.Add("song2","artist2")
	pl.Add("song3","artist3")
	pl.Add("song4","artist4")

	pl.ShowAll()
	pl.StartPlaying()
	pl.NextSong()
	pl.NextSong()
	pl.NextSong()
	pl.NextSong()

}

