package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			select {
			case s := <-c:
				fmt.Println("Выхожу из программы", s)
				os.Exit(0)
			default:
				num := rand.Intn(100)
				square := num * num
				fmt.Println("Квадрат числа", num, "равен", square)
				time.Sleep(time.Second)
			}
		}
	}()

	select {}
}
