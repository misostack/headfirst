package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)

func main() {
	/*
	* - Read votes data from file, input by user
	* - Count names and display the result
	*/
	fmt.Println("[Headfirst Series][Game 11] : Map")
	filePath := os.Args[1:][0]
	fmt.Printf("%v\n", filePath)
	status, err := generateData(filePath)
	if err != nil { fmt.Println(err) }
	if status { fmt.Println("Otoke") }
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