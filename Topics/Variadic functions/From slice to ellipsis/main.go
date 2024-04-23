package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Scan(&num1, &num2, &num3)

	fmt.Println(Evens(num1, num2, num3))
}

func Evens(numbers ...int) bool {
	for _, n := range numbers {
		if n%2 != 0 {
			return false
		}
	}

	return true
}
