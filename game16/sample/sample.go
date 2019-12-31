package main

import (
	"fmt"
	"strconv"
)

type SampleInterface interface {
	MethodWithoutParams()
	MethodWithParams(int64)
	MethodWithReturnValue() string
}

type Sample int64

func (s Sample) MethodWithoutParams() {
	fmt.Println("MethodWithoutParams")
}
func (s Sample) MethodWithParams(num int64) {
	fmt.Printf("MethodWithParams %v + %v = %v\n", s, num, s + Sample(num))
}
func (s Sample) MethodWithReturnValue() string{
	fmt.Println("MethodReturnValue")
	return strconv.Itoa(int(s))
}

func main() {
	// declare an interface
	fmt.Println("Sample interface")
	var s SampleInterface
	s = Sample(5)
	s.MethodWithoutParams()
	s.MethodWithParams(4)
	fmt.Printf("s = %v\n", s.MethodWithReturnValue())	
}