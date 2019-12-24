package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"bufio"
)

var names [10]string = [10]string{ 
	"Dat Xanh Group", 
	"FLC Group", 
	"Vingroup", 
	"CapitaLand",
	"Hung Thinh Real Estate Business Investment Corporation",
	"Novaland",
	"Phat Dat Real Estate Development Corporation",
	"Phu My Hung",
	"Nam Long Investment Corporation",
	"Sun Group",
}

func main() {
	/*
	* - Read votes data from file, input by user
	* - Count names and display the result
	*/
	fmt.Println("[Headfirst Series][Game 11] : Map")
	filePath := os.Args[1:][0]
	fmt.Printf("%v\n", filePath)
	// status, err := generateData(filePath)
	// if err != nil { fmt.Println(err) }
	// if status { fmt.Println("Otoke") }
	var results []int	
	countVotes(filePath, &results)
	printVotesResult(names, results)
}

func printArray(arr []int) {
	for index := 0; index < len(arr); index++ {
		fmt.Printf("%v ", arr[index])
	}
	fmt.Printf("\n")
}

func printVotesResult(names [10]string, results []int) {
	for index := 0; index < len(names); index++ {
		fmt.Printf("%v : %v votes \n", names[index], results[index])
	}
}

func countVotes(filePath string, results *[]int) error{
	// https://www.geeksforgeeks.org/golang-pointer-to-an-array-as-function-argument/
	
	*results = make([]int, 10)
	// reset votes to 0 for all
	for i := 0; i < len(*results); i++ {
		(*results)[i] = 0
	}
	fh, err := os.OpenFile(filePath, os.O_RDONLY, 0700)
	if err != nil {
		fh.Close()
		return err
	}
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		// count votes
		for idx, name := range names {
			if scanner.Text() == name {
				(*results)[idx] ++
				break
			}
		}
	}
	return nil
}

func generateData(filePath string) (status bool, err error) {
	fh, err := os.OpenFile(filePath, os.O_CREATE | os.O_WRONLY, 0700)
	if err != nil {
		return false, err
	}
	var names [10]string = [10]string{ 
		"Dat Xanh Group", 
		"FLC Group", 
		"Vingroup", 
		"CapitaLand",
		"Hung Thinh Real Estate Business Investment Corporation",
		"Novaland",
		"Phat Dat Real Estate Development Corporation",
		"Phu My Hung",
		"Nam Long Investment Corporation",
		"Sun Group",
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	total := int(rand.Int63n(100) + 100)
	
	var content string = ""
	for index := 0; index < total; index++ {
		content += names[rand.Int63n(10)] + "\n"
	}
	// for _, name:= range names {
	// 	content += name + "\n"
	// }
	_, err = fh.WriteString(content)
	// always close file after writing
	fh.Close()	
	if err != nil {
		return false, err
	}
	return true, nil
}