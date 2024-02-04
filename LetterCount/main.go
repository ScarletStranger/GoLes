package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	//fmt.Println("Введите предложение")
	//reader := bufio.NewReader(os.Stdin)
	//text, _ := reader.ReadString('\n')

	text = "phrase txt"

	letters := map[rune]int{
		'q':0,
		'w':0,
		'e':0,
		'r':0,
		't':0,
		'y':0,
		'u':0,
		'i':0,
		'o':0,
		'p':0,
		'a':0,
		's':0,
		'd':0,
		'f':0,
		'g':0,
		'h':0,
		'j':0,
		'k':0,
		'l':0,
		'z':0,
		'x':0,
		'c':0,
		'v':0,
		'b':0,
		'n':0,
		'm':0,
	}

	for _, v := in range unicode.ToLower(text){
		letters[v]++
	}
	
	for a, b := in range letters{
		fmt.Printf("%c: %d\n", a, b)
	}
}
