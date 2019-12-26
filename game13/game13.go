package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	persons := []Person{
		Person{
			ID: 1,
			FirstName: "Nolan",
			LastName: "Nguyen Kenvin",
			DOB: "01/01/1992",
			IsRelatives: false,
			Age: uint8(2019 - 1992),
		},
		Person{
			ID: 2,
			FirstName: "Nalon",
			LastName: "Nguyen Nivnek",
			DOB: "01/01/1995",
			IsRelatives: false,
			Age: uint8(2019 - 1995),
		},		
	}
	err := parseToJsonString(&persons)
	if err != nil { log.Fatal(err) }
	// else
	fmt.Println("OtoKe")
}

func parseToJsonString(data *[]Person) error {
	jsonString, err := json.MarshalIndent(*(data), "","    ")
	if err != nil { log.Fatal(err) }
	return ioutil.WriteFile("./tmp/thedata.json", jsonString, 0644)
}

