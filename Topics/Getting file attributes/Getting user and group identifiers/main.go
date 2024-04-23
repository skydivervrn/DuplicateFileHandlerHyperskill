package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// Do not remove this first line! it gets the Stat of the current directory "."
	fileInfo, err := os.Stat(".")
	if err != nil {
		fmt.Println(err)
	}
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if ok {
		fmt.Println("User identifier:", stat.Uid)
		fmt.Println("Group identifier:", stat.Gid)
	}
}
