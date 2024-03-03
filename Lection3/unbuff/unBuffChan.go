package main

import (
	"fmt"
)

func main() {
	fc := putBook()
	sc := deliverBook(fc)
	tc := burnBook(sc)

	fmt.Println(<-tc)
}

func putBook() chan string {
	firstChan := make(chan string)
	go func() {
		firstChan <- "кладу"
	}()
	return firstChan
}

func deliverBook(firstChan chan string) chan string {
	secondChan := make(chan string)
	fmt.Println(<-firstChan)
	go func() {
		secondChan <- "везу"
	}()
	return secondChan
}

func burnBook(secondChan chan string) chan string {
	thirdChan := make(chan string)
	fmt.Println(<-secondChan)
	go func() {
		thirdChan <- "сжигаю"
	}()
	return thirdChan
}
