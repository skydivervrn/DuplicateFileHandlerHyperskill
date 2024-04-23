package main

import "fmt"

func main() {
	var prefix, name1, name2 string
	fmt.Scan(&prefix, &name1, &name2)

	for _, line := range Greeting(prefix, name1, name2) {
		fmt.Println(line)
	}
}

func Greeting(prefix string, name ...string) []string {
	greetings := make([]string, len(name))
	for i, n := range name {
		greetings[i] = prefix + " " + n
	}
	return greetings
	// your code goes here
}
