package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"os"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(`
		Game rules: 
		- Generate a random number from 0 - 100.
		- Prompt the player to guess what target number is, and store their response
		- If the player's guess is less than target number, says "LOW" else
		says "HIGH"
		- Allow the player to guess up to 10 times. Before each guess let them know how many
		guess they have left
		- If the player's guess is equal to the target number, tell them "Good job". Then stop asking guess
		- If the player ran out of turns without guessing correctly. Says "Sorry, it was [target]\"
	`)
	targetNumber := generateNumber()
	fmt.Printf("The target number is %d\n", targetNumber)
	guessNumber := promptGuessNumber()
	if(guessNumber != -1){
		fmt.Printf("The number %d that you guess is ", guessNumber)	
	}
}

func generateNumber() int{
	// Create and seed the generator.
	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.	
	// fixed seed: rand.Seed(123)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func promptGuessNumber() int64{
	fmt.Print("Please enter your guess number: ")
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')
	// trim space
	input = strings.TrimSpace(input)
	// parse input
	guessNumber, error := strconv.ParseInt(input, 10, 32)
	if error != nil {		
		log.Fatal(error)
		return -1
	}
	// store
	if storeGuessNumber(guessNumber) {
		return guessNumber
	}
	return -1
}

func storeGuessNumber(guessNumber int64) bool{
	fh, err := os.OpenFile("./tmp/guess_number.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
		fh.Close()
		return false
	}
	n, err := fh.WriteString(strconv.FormatInt(guessNumber, 10) + "\n")	
	if n > 0 && err != nil {
		log.Fatal(err)
		fh.Close()
		return false
	}
	fh.Close()
	return true
}