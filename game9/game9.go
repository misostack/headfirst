package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
)

func calculateTotal(numbers ...string) (int, error) {
	if len(numbers) == 0 { return 0, nil }
	total := 0
	for _, num := range numbers {
		n, err := strconv.Atoi(num)
		if err != nil {
			return 0, err
		}
		// else
		total += n
	}
	return total, nil
}

func printList(glue string, strings ...string) string {
	l := len(strings)
	if l == 0 { return "" }
	output := ""
	gl := fmt.Sprintf(" %s ", glue)
	for i, s := range strings {
		if i == l - 1 { gl = "" }
		output += fmt.Sprintf("%s%s", s, gl)
	}
	return output
}

func main() {
	fmt.Println("[Heafirst Series][Game9]: Variadic Function")
	numbers := os.Args[1:]
	total, err := calculateTotal(numbers...)
	if err != nil {
		log.Fatal(err)
	}
	// else
	fmt.Printf("%s = %v\n", printList("+", numbers...), total)
}