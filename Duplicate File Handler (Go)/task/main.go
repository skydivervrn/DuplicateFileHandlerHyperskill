package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var (
	pathsToDelete []string
)

func main() {
	checkArgs()
	folder := os.Args[1]
	fileFormat := ""
	sortingOptions := 0
	checkForDuplicates := false
	//deleteDuplicates := false
	pathsToDelete = make([]string, 0)
	var err error
	for {
		fileFormat = askUserInput("Enter file format:")
		//if fileFormat == "" {
		//	fmt.Println("File format cannot be empty")
		//	continue
		//}
		break
	}
	for {
		sortingOptions, err = strconv.Atoi(askUserInput("Size sorting options:\n1. Descending\n2. Ascending\n"))
		if err != nil {
			log.Fatal(err)
		}
		if sortingOptions < 1 || sortingOptions > 2 {
			fmt.Println("Wrong option")
			continue
		}
		break
	}
	filesMap := searchFiles(folder, fileFormat)
	keys := make([]int64, 0, len(filesMap))
	for k, _ := range filesMap {
		keys = append(keys, k)
	}
	if sortingOptions == 1 {
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] > keys[j]
		})
	}
	if sortingOptions == 2 {
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})
	}
	for _, v := range keys {
		fmt.Printf("%v bytes\n", v)
		for _, pathsArray := range filesMap[v] {
			for _, path := range pathsArray {
				fmt.Println(path)
			}
		}
	}
	for {
		answer := askUserInput("Check for duplicates?")
		if answer == "yes" {
			checkForDuplicates = true
			break
		}
		if answer == "no" {
			break
		}
		fmt.Println("Wrong option")
	}
	if checkForDuplicates {
		for _, v := range keys {
			fmt.Printf("%v bytes\n", v)
			for hash, pathsArray := range filesMap[v] {
				if len(pathsArray) > 1 {
					fmt.Printf("Hash: %x\n", hash)
					for _, path := range pathsArray {
						pathsToDelete = append(pathsToDelete, path)
						fmt.Printf("%v. %s\n", len(pathsToDelete), path)
					}
				}
			}
		}
	}
	for {
		answerDelete := askUserInput("Delete files?")
		if answerDelete == "yes" {
			var fileSize int64
			numbersToDelete := askNumbersToDelete()
			for _, number := range numbersToDelete {
				//num, _ := strconv.Atoi(number)
				fileInfo, err := os.Stat(pathsToDelete[number-1])
				if err != nil {
					fmt.Println(err)
					return
				}
				fileSize += fileInfo.Size()
				os.Remove(pathsToDelete[number-1])
			}
			fmt.Printf("Total freed up space: %v\n", fileSize)
			break
		}
	}
}

func searchFiles(path, fileFormat string) map[int64]map[string][]string {
	var returnMap map[int64]map[string][]string
	returnMap = make(map[int64]map[string][]string)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() {
			return nil
		} else {
			if strings.HasSuffix(info.Name(), fileFormat) {
				file, err := os.Open(path)
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()
				md5Hash := md5.New()
				if _, err = io.Copy(md5Hash, file); err != nil {
					log.Fatal(err)
				}
				if returnMap[info.Size()] == nil {
					returnMap[info.Size()] = map[string][]string{
						string(md5Hash.Sum(nil)): {},
					}
				}
				returnMap[info.Size()][string(md5Hash.Sum(nil))] = append(returnMap[info.Size()][string(md5Hash.Sum(nil))], path)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return returnMap
}

func checkArgs() {
	if len(os.Args) < 2 {
		fmt.Println("Directory is not specified")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	}
}

func askUserInput(str string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(str)
	buf, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	return string(buf)
}

func askNumbersToDelete() (numbersToDeleteArray []int) {
	for {
		numbersToDelete := askUserInput("Enter file numbers to delete:")
		if numbersToDelete == "" {
			fmt.Println("Wrong format")
			continue
		}
		for _, number := range strings.Split(numbersToDelete, " ") {
			numberToCheck, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println("Wrong format")
				continue
			}
			if numberToCheck < 1 || numberToCheck > len(pathsToDelete) {
				fmt.Println("Wrong format")
				continue
			}
			num, _ := strconv.Atoi(number)
			numbersToDeleteArray = append(numbersToDeleteArray, num)
		}
		return
	}
}
