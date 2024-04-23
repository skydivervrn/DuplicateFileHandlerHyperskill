package main

import (
	"fmt"
	"log"
	"os"
)

// nolint: gomnd // <-- DO NOT delete this line!
func main() {
	// DO NOT modify the contents of the `data` variable!
	data := `"One","Metallica","...And Justice for All",7:27
"Fuel","Metallica","Reload",4:30
"Master of Puppets","Metallica","Master of Puppets",8:36
`
	// Write the code to append the additional records to the "songs.csv" file below
	// Use the os.OpenFile() function first to open "songs.csv" with the os.O_APPEND and os.O_WRONLY flags
	// And then you can use the fmt.Fprintln() function to write the data to the file.
	file, err := os.OpenFile("songs.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Print(err)
	}
	fmt.Fprintln(file, data)
}
