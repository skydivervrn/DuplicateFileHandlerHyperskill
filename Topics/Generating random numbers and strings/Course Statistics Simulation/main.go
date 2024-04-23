package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var seed int64
	fmt.Scanln(&seed)
	rand.Seed(seed)

	var arr [10]float64
	var sum float32
	for i := 0; i < 10; i++ {
		arr[i] = rand.Float64() * 100
		sum += float32(arr[i])
	}

	fmt.Printf("%.2f\n", sum/10)
}
