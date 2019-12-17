package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	// len : https://www.dotnetperls.com/len-go
	fmt.Println("Please enter your username:")
	reader := bufio.NewReader(os.Stdin)
	username, error := reader.ReadString('\n')
	if(error != nil) {fmt.Println(error)}
	if(error == nil && len(username) > 0) {fmt.Println("Username:", username)}
}