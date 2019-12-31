package main

import (
	"fmt"
)

type Player interface {
	Play(string)
	Stop()
}

type TapePlayer struct {
	Title string
	Batteries string
}
func (t TapePlayer) Play(song string) {
	fmt.Printf("%v : Playing %v\n", t.Title, song)
}
func (t TapePlayer) Stop() {
	fmt.Printf("%v : Stop playing\n", t.Title)
}

type TapeRecorder struct {
	Title string
	Microphones int
}
func (t TapeRecorder) Play(song string) {
	fmt.Printf("%v : Playing %v\n", t.Title, song)
}
func (t TapeRecorder) Stop() {
	fmt.Printf("%v : Stop playing\n", t.Title)
}
func (t TapeRecorder) Recording() {
	fmt.Printf("%v : Recording\n", t.Title)
}

func playList(device Player, songs []string) {	
	for _,s:= range songs {
		device.Play(s)
	}
	device.Stop()
}

func main() {
	fmt.Println("[Headfirst Series][Game 16: Interface]")
	mixtape := []string{ "Hello", "ABC", "DEF" }
	tp := TapePlayer{"Tape Player",""}
	tr := TapeRecorder{"Tape Recorder", 0}
	playList(&tp, mixtape)
	playList(&tr, mixtape)
}