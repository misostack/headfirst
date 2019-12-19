package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("[Headfirst Series][Game 8]: The arrays")
	var notes [7]string
	for i:=0; i < len(notes); i++ {
		notes[i] = strconv.Itoa(i + 1)
		fmt.Println(notes[i])
	}
	var numbers [10]int
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
	tn := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	it := 0
	for true {
		fmt.Println(tn[it])
		it++
		if it == 10 { break }
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
}

func sum(arr [10]int) int {
	s:= 0
	for _, ele := range arr {
		s += ele
	}
	return s
}