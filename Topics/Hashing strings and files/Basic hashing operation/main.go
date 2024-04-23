package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func main() {
	// DO NOT delete this code block! it takes as an input the famous movie password
	var famousPwd string
	fmt.Scanln(&famousPwd)

	// Use the correct function to create 'new' MD5 Hash and SHA-256 Hash interfaces below:
	md5Hash := md5.New()
	sha256Hash := sha256.New()

	// Call the correct method that adds the 'famousPwd' to the hash that will be calculated:
	md5Hash.Write([]byte(famousPwd))
	sha256Hash.Write([]byte(famousPwd))

	// Call the method that returns the computed hash slice
	// And print the hash in hexadecimal notation below:
	fmt.Printf("%x\n", md5Hash.Sum(nil))
	fmt.Printf("%x\n", sha256Hash.Sum(nil))
}
