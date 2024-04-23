package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sample.txt") // open 'sample.txt' here with the os.Open function
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lineCount int // lineCount will be used to count the number of lines in the file

	scanner := bufio.NewScanner(file) // create a new scanner for the file
	scanner.Split(bufio.ScanLines)
	// Use the for scanner.Scan() loop to read the file line by line
	// and increment the lineCount variable each loop
	for scanner.Scan() {
		lineCount++
	}

	fmt.Println(lineCount) // print the total lineCount here!
}
