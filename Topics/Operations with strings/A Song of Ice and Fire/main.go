package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	westeros = "Westeros"
	misteros = "Misteros"
)

// Which function from the `strings` package is used to replace a string?
func wordReplacer(novel, wrong, correct string) string {
	return strings.Replace(novel, wrong, correct, -1)
}

// DO NOT delete or modify the contents of the main function!
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(wordReplacer(scanner.Text(), misteros, westeros))
}
