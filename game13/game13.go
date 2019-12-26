package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"bufio"
	"os"
	"strings"
	"strconv"
	"time"
)

const (
	RELATIVES_STRING = "(relatives)"
)

type Person struct {
	ID uint32
	FirstName string
	LastName string
	IsRelatives bool
	DOB string
	Age uint8
}

func main() {
	fmt.Println("[Headfirst Series] Game 13 : The Challenger, read the data and transform it to json file")
	// data structure
	// xx	ABC DEF XYZ (relatives)	DD/MM/YYYY
	// Todo
	// - Read the file
	// - Split the data in a line into a struct {id, name, isRelatives?, dob}
	// - Save new data into a new file with json format
	lines, err := readLines("./tmp/thedata.txt")
	persons := []Person{}
	if err != nil { log.Fatal(err) }
	for _, line := range lines {
		p := Person{} // empty struct
		convertLineToPerson(line, &p)
		persons = append(persons, p)
	}
	
	// persons := []Person{
	// 	Person{
	// 		ID: 1,
	// 		FirstName: "Nolan",
	// 		LastName: "Nguyen Kenvin",
	// 		DOB: "01/01/1992",
	// 		IsRelatives: false,
	// 		Age: uint8(2019 - 1992),
	// 	},
	// 	Person{
	// 		ID: 2,
	// 		FirstName: "Nalon",
	// 		LastName: "Nguyen Nivnek",
	// 		DOB: "01/01/1995",
	// 		IsRelatives: false,
	// 		Age: uint8(2019 - 1995),
	// 	},		
	// }
	err = parseToJSONString(&persons)
	if err != nil { log.Fatal(err) }
	fmt.Printf("Write file completed %v\n", "./tmp/thedata.json")

	// l := "60	Nguyễn Thế Vinh (relatives)	13/02/2013"
	// parts := strings.Split(l, "\t")
	// id, name, dob := parts[0], parts[1], parts[2]
	// fmt.Printf("%v.\t%v\t%v\n", id, name, dob)	
}

func parseToJSONString(data *[]Person) error {
	jsonString, err := json.MarshalIndent(*(data), "","    ")
	if err != nil { log.Fatal(err) }
	return ioutil.WriteFile("./tmp/thedata.json", jsonString, 0644)
}

func calculateAgeFromDOB(s string) uint8 {
	currentYear, err := strconv.Atoi(strings.Split(s, "/")[2])
	if err != nil { log.Fatal(err) }	
	return uint8(time.Now().Year() - currentYear)
}

func convertLineToPerson(l string, p *Person)  {
	// data structure
	// xx	ABC DEF XYZ (relatives)	DD/MM/YYYY	
	// - Split the data in a line into a struct {id, name, isRelatives?, dob}
	parts := strings.Split(l, "\t")
	id, _ := strconv.Atoi(parts[0])
	name, dob := parts[1], parts[2]
	IsRelatives := false
	if strings.Contains(name, RELATIVES_STRING) {
		IsRelatives = true
	}
	name = strings.TrimSpace(strings.Replace(name, RELATIVES_STRING, "", -1))
	names := strings.Split(name, " ")
	lastname := names[0:len(names) - 1]
	*p = Person{
		ID: uint32(id),
		FirstName: names[len(names) - 1],
		LastName: strings.Join(lastname, " "),
		DOB: dob,
		Age: calculateAgeFromDOB(dob),
		IsRelatives: IsRelatives,
	}
}

func readLines(filepath string) ([]string, error) {
	// reader := bufio.NewReader()
	lines := []string{}
	fh, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		fh.Close()
		return nil, nil
	}
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fh.Close()
	return lines, nil
}