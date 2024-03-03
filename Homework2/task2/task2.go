package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("task2.txt")
	if err != nil {
		fmt.Println("Не удалось открыть файл", err)
		return
	}
	defer file.Close()
	buffer := make([]byte, 8)
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(fileInfo.Size())
	if fileInfo.Size() == 0 {
		fmt.Println("Файл пуст")
		return
	}
	if _, err := io.ReadFull(file, buffer); err != nil {
		fmt.Println("Не удалось прочитать файл", err)
		return
	}
	fmt.Printf("%s\n", buffer)
}
