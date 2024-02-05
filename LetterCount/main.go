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
	letters := map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
		'h': 0,
		'i': 0,
		'j': 0,
		'k': 0,
		'l': 0,
		'm': 0,
		'n': 0,
		'o': 0,
		'p': 0,
		'q': 0,
		'r': 0,
		's': 0,
		't': 0,
		'u': 0,
		'v': 0,
		'w': 0,
		'x': 0,
		'y': 0,
		'z': 0,
	}

	for _, v := range strings.ToLower(strings.ReplaceAll(text, " ", "")) {
		if unicode.IsLetter(v) {
			letters[v]++
			letterCount++
		}
	}

	for a, b := range letters {
		if b > 0 {
			fmt.Printf("%c: %d %f\n", a, b, float32(letterCount/letters[b]))
		}
	}
}
