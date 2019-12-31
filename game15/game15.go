package main

import (
	"fmt"	
	"log"
	"github.com/misostack/headfirst/game15/calendar"
)

func main() {
	fmt.Println("[Headfirst Series]: Game 15 - Encapsulation and Embedding")	
	// date := calendar.Date {}
	// err := date.SetYear(1) // automatic convert date to a pointer
	// if err != nil { log.Fatal(err) }
	// err = date.SetMonth(2)
	// if err != nil { log.Fatal(err) }
	// err = date.SetDay(28)
	// if err != nil { log.Fatal(err) }
	// fmt.Printf("%v/%v/%v\n", date.Day(), date.Month(), date.Year())
	event := calendar.Event {}	
	err := event.SetTitle("New Year 2020")
	if err != nil { log.Fatal(err) }
	event.SetYear(2020)
	event.SetMonth(1)
	event.SetDay(1)
	fmt.Printf("\"%v\" will be happened on %v/%v/%v\n", event.Title(), event.Day(), event.Month(), event.Year())
}