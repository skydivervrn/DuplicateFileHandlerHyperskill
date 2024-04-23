package main

func main() {
	// Create a set named 'mySet' with an empty struct as the value type below.
	// The set can have as many keys as you want; you can get creative!
	mySet := map[string]struct{}{
		"Cool": struct{}{},
	}

	checkSetTypeValues(mySet) // DO NOT delete this line!
}

//func checkSetTypeValues(strasd map[string]struct{}) {
//	fmt.Print(strasd)
//}
