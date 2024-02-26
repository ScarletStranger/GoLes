package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("task4.txt")
	if err != nil {
		fmt.Println("Не удалось создать файл", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	args := strings.Fields(text)
	er := os.WriteFile(file.Name(), []byte(strings.Join(args, " ")), 0666)
	if er != nil {
		fmt.Println("Не удалось записать данные в файл", err)
	}
	output, _ := os.ReadFile(file.Name())
	fmt.Println(string(output))
}
