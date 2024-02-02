package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	var fileName, fileExt string

	filePath := os.Args[1]
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fileName = fileInfo.Name()
	lastIndex := len(fileName) - 1
	for i := len(fileName) - 1; i >= 0; i-- {
		if fileName[i] == '.' {
			lastIndex = i
			break
		}
	}

	fileExt = fileName[lastIndex+1:]
	fileName = fileName[:lastIndex]

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
