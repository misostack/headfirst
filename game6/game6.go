package main

import (
	"fmt"
	"reflect"
	"strings"
)
/*
* Explaination: The function required a string pointer
* The function will transform the value at this pointer to uppercase
* and update it with new value
* 
*/
func toUpperCase(str *string) {
	*str = strings.ToUpper(*str)
}

func main() {
	fmt.Println("[Headfirst Series][Game 6]: Pointers in Golang!")
	x := 10
	fmt.Printf("x value: %v\n", x)
	fmt.Printf("x address: %v\n", &x)
	fmt.Printf("x value type: %v\n", reflect.TypeOf(x))
	fmt.Printf("x address type: %v\n", reflect.TypeOf(&x))
	xPointer := &x
	// The aterisk operator
	*xPointer = 8
	fmt.Printf("xPointer address: %v\n", &xPointer)
	fmt.Printf("xPointer value: %v\n", *xPointer)
	fmt.Printf("x value: %v\n", x)
	// toUpperCase
	aString := "the lower case string"
	toUpperCase(&aString)
	fmt.Println(aString)
}