package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	var value int
	fmt.Println("Введите число")
	if _, err := fmt.Fscan(os.Stdin, &value); err != nil {
		log.Fatal(err)
	}

	if isPrime(value) {
		fmt.Println("Число", value, "простое")
	} else {
		fmt.Println("Число", value, "не простое")
	}
}
