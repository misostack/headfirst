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
	// fmt.Printf("The target number is %d\n", targetNumber)
	guessTimes := 0
	guessTimesLimit := 10
	guessTimesText := "turn"
	guessCompareString := ""
	for true {
		guessTimes++
		if guessTimes > guessTimesLimit { 
			fmt.Printf(
				"Sorry, the target is %d",
				targetNumber,
			)
			break
		}
		guessNumber := promptGuessNumber(guessTimes)
		if guessNumber == targetNumber {
			if guessTimes > 1 { guessTimesText = "turns" }
			fmt.Printf(
				"Good job! The target number is %d , and you take %d %s to find out!\n",
				targetNumber,
				guessTimes,
				guessTimesText,
			)
			break
		} else {
			if guessNumber > targetNumber {
				guessCompareString = "HIGHER"
			} else {
				guessCompareString = "LOWER"
			}
			fmt.Printf("Your guess number is %s than the target number\n", guessCompareString)
		}
	}
	
}

func generateNumber() int64{
	// Create and seed the generator.
	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.	
	// fixed seed: rand.Seed(123)
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(100)
}

func promptGuessNumber(guessTimes int) int64{
	guessTimesText := "time"
	if guessTimes > 1 { guessTimesText = "times" } 
	fmt.Printf("Please enter your guess number (%d %s): ", guessTimes, guessTimesText)
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