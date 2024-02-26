package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("task3.txt")
	if err = os.Chmod("task3.txt", 0444); err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println("Не удалось записать данные в файл файл", err)
		return
	}
	defer file.Close()
	writer.WriteString("test")
	writer.Flush()
}
