package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("task1.txt")
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println("Не удалось создать файл", err)
		return
	}
	defer file.Close()
	var text string
	var date = time.Now()
	for {
		if text == "exit" {
			break
		}
		fmt.Println("Введите сообщение")
		fmt.Scan(&text)
		writer.WriteString(date.Format("2006-01-02"))
		writer.WriteString(date.Format("15:04:05"))
		writer.WriteString(" ")
		writer.WriteString(text)
		writer.WriteString("\n")
		writer.Flush()
	}
}
