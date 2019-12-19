package tools

import (
	"bufio"
	"os"
	"strconv"
)

// ReadFile will require a filepath
// It will read the file content and parse the data into array
func ReadFile(filepath string) ([10]int, error){
	fh, err := os.Open(filepath)
	var arr [10]int
	if err != nil {
		fh.Close()
		return arr, err
	}
	// else
	scanner := bufio.NewScanner(fh)
	i := 0
	for scanner.Scan() {
		arr[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			return arr, err
		}
		i++
		if i == 10 { break }
	}	
	err = fh.Close()
	if err != nil {
		return arr, err
	}	
	if scanner.Err() != nil {
		return arr, scanner.Err()
	}			
	return arr, nil
}