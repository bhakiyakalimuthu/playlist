package linkedlist

import "fmt"

type SongS struct {
	name string
	artist string
	next *SongS
}

type PlayListS struct{
	name string
	head *SongS
	nowPlaying *SongS
}
func NewPlayListS(name string) *PlayListS {
	return &PlayListS{
		name:       name,
	}
}

func (p *PlayListS) Add(name,artist string)  {
	song := &SongS{
		name:   name,
		artist: artist,
	}
	if p.head == nil {
		p.head = song
	} else {
		current := p.head
		for current.next != nil  {
			current = current.next
		}
		current.next = song
	}
}

func (p *PlayListS) ShowAll()  error{
	if p.head == nil {
		return fmt.Errorf("empty playlist")
	}
	current := p.head
	fmt.Printf("song - %s %s \n", current.artist,current.name)
	if current.next != nil {
		fmt.Printf("song - %s %s \n", current.next.artist,current.next.name)
		current = current.next
	}
	fmt.Printf("song - %s %s \n", current.next.artist,current.next.name)
	return nil
}

func (p *PlayListS) StartPlaying()  *SongS{
	p.nowPlaying = p.head
	fmt.Printf("start playing - %s %s \n", p.nowPlaying.artist,p.nowPlaying.name)
	return p.nowPlaying
}

func (p *PlayListS) NextSong()  *SongS{
	if p.nowPlaying.next == nil {
		fmt.Printf("no next song")
		return nil
	}
	p.nowPlaying = p.nowPlaying.next
	fmt.Printf("next song - %s %s \n", p.nowPlaying.artist,p.nowPlaying.name)
	return p.nowPlaying
}

func (p *PlayListS) PreviousSong()  {
// 	not possible in singly linked list
}
