package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"bufio"
	"log"
	"sort"
)


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
	// results := []int{}
	// names := []string{}
	// countVotes(filePath, &names, &results)
	// printVotesResult(names, results)
	ranks, err := countVotesWithMap(filePath)
	if err != nil { log.Fatal(err) }
	names := []string{}
	for name := range ranks {
		names = append(names, name)
	}
	sort.Strings(names)
	for idx, name := range names {
		fmt.Printf("%v. %v : %v\n", idx + 1, name, ranks[name])
	}
	// zero values
	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)
	// delete
	delete(counters, "a")
	value, ok = counters["a"]
	fmt.Println(value, ok)		
}

func printArray(arr []int) {
	for index := 0; index < len(arr); index++ {
		fmt.Printf("%v ", arr[index])
	}
	fmt.Printf("\n")
}

func printVotesResult(names []string, results []int) {
	for index := 0; index < len(names); index++ {
		fmt.Printf("%v. %v : %v votes \n", index + 1, names[index], results[index])
	}
}

func readLines(filePath string) ([]string, error) {
	lines := []string{} // empty slice
	fh, err := os.OpenFile(filePath, os.O_RDONLY, 0700)
	if err != nil {
		fh.Close()
		return lines, err
	}
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fh.Close()
	return lines, nil
}

func countVotesWithMap(filePath string) (map[string]int, error) {
	lines, err := readLines(filePath)
	if err != nil { return nil, err }
	var ranks map[string]int // key:string, value:int	
	ranks = make(map[string]int)
	for _,line := range lines {
		ranks[line]++
	}
	fmt.Println(ranks)
	// ranks["gold"] = 1
	// ranks["silver"] = 2
	// fmt.Printf("ranks[\"gold\"]: %v\n", ranks["gold"])
	// fmt.Println(lines[0])
	// The for ... range loop processes map keys and values in a random order because a map is
	// an unordered collection of keys and values
	
	return ranks, nil
}

func countVotes(filePath string, names *[]string, counts *[]int) error{
	// https://www.geeksforgeeks.org/golang-pointer-to-an-array-as-function-argument/

	
	lines, err := readLines(filePath)
	if err != nil { return err }
	for _, line := range lines {
		// count votes
		matched := false
		for idx, name := range (*names) {
			if line == name {
				matched = true
				(*counts)[idx]++
				break
			}
		}
		if !matched {			
			*names = append(*names, line)
			*counts = append(*counts, 1)
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