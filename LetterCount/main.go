package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.ReadFile("phrase.txt")
	if err != nil {
		fmt.Print(err)
	}

	text := string(file)
	letterCount := 0
	letters := map[rune]int{}

	for _, v := range strings.ToLower(strings.ReplaceAll(text, " ", "")) {
		if unicode.IsLetter(v) {
			letters[v]++
			letterCount++
		}
	}

	for key, value := range letters {
		fmt.Printf("%c - %d %.2f\n", key, value, float32(value/letterCount))
	}
}
