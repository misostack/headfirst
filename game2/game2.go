package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"log"
	s "strings"
)

func main() {
	fmt.Print("Please enter any number : ")
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')
	if error != nil {
		log.Fatal(error)
		return
	}
	// trim space
	input = s.TrimSpace(input)
	// parse input
	number, error := strconv.ParseInt(input, 10, 32)
	if error != nil {		
		log.Fatal(error)
		return
	}
	// check
	if number % 2 == 0 {
		fmt.Println(number, " is even number")		
	} else {
		fmt.Println(number, " is odd number")
	}
}