package main

import (
	"fmt"
	"strconv"
	"log"
	"os"
	game8tools "github.com/misostack/headfirst/game8/tools"
)

func main() {	
	fmt.Println("[Headfirst Series][Game 8]: The arrays")
	args := os.Args[1:]
	fmt.Println("Args:")
	for _ , arg := range args {
		fmt.Printf("%v", arg)
	}
	fmt.Println("------------")
	var notes [7]string
	for i:=0; i < len(notes); i++ {
		notes[i] = strconv.Itoa(i + 1)
		fmt.Println(notes[i])
	}
	numbers := make([]int, 10)
	for i:=0; i < len(numbers); i++ {
		fmt.Printf("%d ", numbers[i])
		numbers[i] = i + 1
		fmt.Printf("%d ", numbers[i])
	}
	fmt.Printf("\n")

	var mnotes [7]string = [7]string{ "do", "re", "mi", "fa", "so", "la", "ti" }
	for i:=0; i<len(mnotes); i++ {
		fmt.Printf("%s ",mnotes[i])
		if i == len(mnotes) - 1 { fmt.Println("")}
	}
	tn := numbers[2:5]
	// test append a slice is underlying an array
	tn = append(tn, 100) // 
	it := 0
	fmt.Println("SLICE FROM 2:5")
	for true {
		fmt.Println(tn[it])
		it++
		if it == len(tn) { break }
	}	

	// another way to loop
	for idx, mnote := range mnotes {
		fmt.Printf("idx: %v, value: %v\n", idx, mnote)
	}

	for _, num := range numbers {
		fmt.Printf("%v + ", num)		
	}
	// sum
	fmt.Printf("= %v\n", sum(numbers))

	// read from file
	data, err := game8tools.ReadFile("./tmp/game8.txt")
	if err != nil {
		log.Fatal(err)
	}
	// else
	fmt.Printf("= %v\n", sum(data))
}

func sum(arr []int) int {
	s:= 0
	for _, ele := range arr {
		s += ele
	}
	return s
}