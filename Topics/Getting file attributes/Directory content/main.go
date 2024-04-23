package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    // DO NOT delete the code line below! It returns a slice with the list of files:
    fileNames := getFileNames()

    for _, fileName := range fileNames {
        fileInfo, ? := ?
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%s %5t %s\n", ?, ?, ?)
    }
}