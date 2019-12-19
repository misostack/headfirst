package main

import (
	"fmt"
	// "errors"
	"log"
)

func theSquareOfTreeAngles(a int, b int) (square int, err error ){
	if a < 0 {
		return 0, fmt.Errorf("a with value %v is invalid", a)
	}
	if b < 0 {
		return 0, fmt.Errorf("b with value %v is invalid", b)
	}
	return a * b, nil
}

func main() {
	fmt.Println("[Headfirst Series][Game 5]: Raise errors")
	// err := errors.New("Something went wrong")
	// log.Fatal(err)
	a1, b1 := -1, 5
	s1, err1 := theSquareOfTreeAngles(a1, b1)
	if err1 != nil {
		log.Fatal(err1)
		return
	}
	// else
	fmt.Printf("a=%d and b=%d ==> s=%d", a1, b1, s1)
}