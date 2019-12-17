package main

import (
	"fmt"
	"math"
	"strconv"
	"os"
	"reflect"
	"time"
	"bufio"	
)

const aUPPCASE, aLOWERCASE, zUPPERCASE, zLOWERCASE = 'A', 'a', 'Z', 'z'

func main() {

	fmt.Println("Please enter your name")
	reader := bufio.NewReader(os.Stdin)
	username, error := reader.ReadString('\n')
	if(error != nil){
		fmt.Println(error)
	}else{
		fmt.Println("Username:", username)
	}
	

	a := "An example"

	fmt.Println("Headfirst GO")
	fmt.Println('A', string('A'))
	fmt.Println('Z')
	fmt.Println("The alphabet")	
	theAlphabet(true)
	theAlphabet(false)
	theConversion()
	theEvenNumbers(1, 10)
	theOddNumbers(1, 10)
	binaryFormat(10)
	theSQRT(4)
	fmt.Fprintf(os.Stdout, "System print something\n")
	fmt.Fprintf(os.Stderr, "an %s!\n", "error")
	fmt.Println(reflect.TypeOf('A'))
	fmt.Println(reflect.TypeOf("A"))
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(1.2))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(0xFF))
	fmt.Println(a)
	theCurrentTime()
}

func theAlphabet(uppercase bool) {
	var str = ""
	if uppercase {
		for i:= aUPPCASE; i<= zUPPERCASE; i++ {
			str += string(i) + " "
		}
	}else {
		for i:= aLOWERCASE; i<= zLOWERCASE; i++ {
			str += string(i) + " "
		}		
	}
	fmt.Println(str)
}

func theConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
}

func theEvenNumbers(start int, end int) {
	theNumbers(start, end, false)
}

func theOddNumbers(start int, end int) {
	theNumbers(start, end, true)
}

func theNumbers(start int, end int, isOdd bool) {
	var str = "\n"
	if isOdd { str += "The odd numbers:\n" }
	if !isOdd{ str += "The even numbers:\n" }

	for i:= start; i<= end; i++ {
		if(isOdd && i % 2 != 0) {
			str += strconv.Itoa(i) + " "
		}
		if(!isOdd && i % 2 == 0) {
			str += strconv.Itoa(i) + " "
		}
	}
	fmt.Println(str)
}

func binaryFormat(number uint) string{
	// fmt.Printf("Binary format of %d:\n", number)
	// fmt.Printf("%b\n", number)
	return fmt.Sprintf("%b", number)
}

func theSQRT(number uint) {
	len := 0
	for true {		
		len++
		fmt.Printf("%s >> %d = %d\n", binaryFormat(number), len, number >> len)
		if(len == 2){			
			break
		}
	}
}

func theCurrentTime() {
	now := time.Now()
	var month, year = now.Month(), now.Year()
	var dayFunc = now.Day
	
	fmt.Println(now)
	fmt.Println(dayFunc(), month, year)
}