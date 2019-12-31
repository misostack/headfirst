package main

import (
	"fmt"
)

type TapePlayer struct {
	Batteries string
}
func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stop")
}

type TapeRecorder struct {
	Microphones int
}
func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapeRecorder) Stop() {
	fmt.Println("Stop")
}
func (t TapeRecorder) Recording() {
	fmt.Println("Recording")
}

func playList(device *TapePlayer, songs []string) {	
	for _,s:= range songs {
		device.Play(s)
	}
	device.Stop()
}

func main() {
	fmt.Println("[Headfirst Series][Game 16: Interface]")
	mixtape := []string{ "Hello", "ABC", "DEF" }
	tp := TapePlayer{}
	tr := TapeRecorder{}
	playList(&tp, mixtape)
	playList(&tr, mixtape)
}