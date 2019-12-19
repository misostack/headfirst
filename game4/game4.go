package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("[Headfirst Series][Game 4][Formatting Verbs]")
	// numbers
	a, b := 6, 7
	c := math.Sqrt(float64(a * a + b * b)) 
	fmt.Printf("The 3 angles has a=%d and b=%d => c=%0.2f\n", a, b, c)
	// strings
	fmt.Printf("%s\n", "This is format string!")
	// boolean
	fmt.Printf("%d > %d ? %t\n", a, b, a > b)
	fmt.Printf("%d/%d = %v\n", a, b, a / b)
	fmt.Printf("%d/%d = %#v\n", a, b, "\n\t")
	fmt.Printf("%d/%d = %T\n", a, b, a / b)
	fmt.Printf("%d/%d = %0.2f%%\n", a, b, float64(a)/float64(b)*100)
	// draw tables
	fmt.Printf("%12s | %s\n", "Product", "Price in euros")
	// random
	rand.Seed(time.Now().UnixNano())	
	i:=0
	for true {
		i++
		fmt.Printf("%12s | %2d euros\n", "Product " + strconv.Itoa(i), rand.Int63n(98)+1)
		if i > 10 { break }	
	}
}