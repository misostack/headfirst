package main

import (
	"fmt"
	"log"
)

type Player interface {
	Play(string)
	Stop()
	Title() string
}

type TapePlayer struct {
	title string
	Batteries string
}
func (t TapePlayer) Play(song string) {
	fmt.Printf("%v : Playing %v\n", t.title, song)
}
func (t TapePlayer) Stop() {
	fmt.Printf("%v : Stop playing\n", t.title)
}
func (t TapePlayer) Title() string{
	return t.title
}

type TapeRecorder struct {
	title string
	Microphones int
}
func (t TapeRecorder) Play(song string) {
	fmt.Printf("%v : Playing %v\n", t.title, song)
}
func (t TapeRecorder) Stop() {
	fmt.Printf("%v : Stop playing\n", t.title)
}
func (t TapeRecorder) Recording() {
	fmt.Printf("%v : Recording\n", t.title)
}
func (t TapeRecorder) Title() string{
	return t.title
}

func playList(device Player, songs []string) {	
	for _,s:= range songs {
		device.Play(s)
	}
	device.Stop()
}

func tryOut(d Player, song string) {
	d.Play(song)
	d.Stop()
	recorder, ok := d.(TapeRecorder)
	if ok {
		recorder.Recording()
	} else {
		fmt.Printf("%v was not a TapeRecorder.\n", d.Title())
	}
}

type BadRequestError string

func (b BadRequestError) Error() string{
	return "BadRequest:" + string(b)
}

func accessBackOffice(usertype string) error{
	if usertype != "admin" {
		return BadRequestError("You don't have permissions to access this page")
	}
	return nil
}

func main() {
	fmt.Println("[Headfirst Series][Game 16: Interface]")
	// mixtape := []string{ "Hello", "ABC", "DEF" }
	// tp := TapePlayer{"Tape Player",""}
	// tr := TapeRecorder{"Tape Recorder", 0}
	// tr.Recording()
	// playList(&tp, mixtape)
	// playList(&tr, mixtape)
	tryOut(TapeRecorder{title:"TapeAAA"}, "Happy New Year")
	tryOut(TapePlayer{title:"TapePlayerBBB"}, "hello")
	err := accessBackOffice("user")
	if err != nil {
		log.Fatal(err)
	}
	// else
	fmt.Println("You've just accessed to dashboard!")
}