package main

import (
	"fmt"
	"math/rand"
)

func main() {
	upperCase := "ABCDEFGHIJKLMOPQRSTUVWXYZ"
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	number := "0123456789"
	specialSymbol := "!?$&%#"

	var seed int64
	fmt.Scanln(&seed)
	rand.Seed(seed)

	fmt.Println(
		string(upperCase[rand.Intn(len(upperCase))]) +
			string(lowerCase[rand.Intn(len(lowerCase))]) +
			string(number[rand.Intn(len(number))]) +
			string(specialSymbol[rand.Intn(len(specialSymbol))]),
	)
}
