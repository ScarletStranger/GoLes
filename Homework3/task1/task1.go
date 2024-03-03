package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	for {
		fmt.Println("Введите 'стоп' для остановки программы\nВведите число:")
		var input string
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}
		if input == "стоп" {
			fmt.Println("Остановка программы...")
			os.Exit(0)
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка преобразования:", err)
			continue
		}
		var wg sync.WaitGroup
		wg.Add(2)
		go multiply(&wg, square(&wg, number))
		wg.Wait()
	}
}

func square(wg *sync.WaitGroup, number int) int {
	defer wg.Done()
	number *= number
	fmt.Println("Квадрат:", number)
	return number
}

func multiply(wg *sync.WaitGroup, number int) {
	defer wg.Done()
	number *= 2
	fmt.Println("Умножение на 2:", number)
}
