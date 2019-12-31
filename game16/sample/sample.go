package main

import (
	"fmt"
	"strconv"
)

type SampleInterface interface {
	MethodWithoutParams()
	MethodWithParams(int64)
	MethodWithReturnValue() string
}

type Sample int64

func (s Sample) MethodWithoutParams() {
	fmt.Println("MethodWithoutParams")
}
func (s Sample) MethodWithParams(num int64) {
	fmt.Printf("MethodWithParams %v + %v = %v\n", s, num, s + Sample(num))
}
func (s Sample) MethodWithReturnValue() string{
	fmt.Println("MethodReturnValue")
	return strconv.Itoa(int(s))
}

type NoiseMaker interface {
	makeNoise()
}

type Appliance interface {
	turnOn()
}

type Whistle string
type Horn string
type Robot string
type Fan string
type CoffeePot string

func (w Whistle) makeNoise() {
	fmt.Printf("%v : tutu\n", w)
}

func (h Horn) makeNoise() {
	fmt.Printf("%v : toetoe\n", h)
}

func playMusicInstrument(n NoiseMaker) {
	n.makeNoise()	
}

func (r Robot) makeNoise() {
	fmt.Printf("%v : pi pi pi\n", r)
}

func (r Robot) walk() {
	fmt.Println("%v : is walking slowly\n", r)
}

func (f Fan) turnOn() {
	fmt.Printf("%v : Spinning\n", f)
}

func (c CoffeePot) turnOn() {
	fmt.Printf("%v : Powering up\n", c)
}

func (c CoffeePot) brew() {
	fmt.Printf("%v : Heating up\n", c)
}

func (c CoffeePot) String() string {
	return "I'm a little coffee pot!"
}

func main() {
	// declare an interface
	fmt.Println("Sample interface")
	var s SampleInterface
	s = Sample(5)
	s.MethodWithoutParams()
	s.MethodWithParams(4)
	fmt.Printf("s = %v\n", s.MethodWithReturnValue())
	var toy NoiseMaker
	toy = Whistle("whistle")
	playMusicInstrument(toy)
	toy = Horn("horn")
	playMusicInstrument(toy)
	toy = Robot("robot")
	playMusicInstrument(toy)
	var appliance Appliance
	appliance = Fan("Fan")
	appliance.turnOn()
	appliance = CoffeePot("Coffee Pot")
	appliance.turnOn()
	coffeebot, ok := appliance.(CoffeePot)
	if ok {
		fmt.Println(coffeebot.String())
	}
}