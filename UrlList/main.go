package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Item struct {
	Name string
	Date time.Time
	Tags string
	Link string
}

func main() {
	defer func() {

		_ = keyboard.Close()
	}()
	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода и приложения нажмите Esc")

OuterLoop:

	for {

		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		urls := make([]Item, 0)
		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильный аргументы в формате url описание теги")
				continue OuterLoop
			}

			url := Item{
				Name: args[1],
				Date: time.Now(),
				Tags: args[2],
				Link: args[0],
			}
			urls = append(urls, url)

		case 'l':
			for idx := range urls {
				fmt.Println("Имя: ", urls[idx].Name, `\n`,
					"URL:", urls[idx].Link, `\n`,
					"Теги:", urls[idx].Tags, `\n`,
					"Дата:", urls[idx].Date, `\n`)
			}

		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Введите имя ссылки, которое нужно удалить")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			_ = text

			for idx := range urls {
				if text == urls[idx].Name {
					urls = append(urls[:idx], urls[idx+1:]...)
				}
			}
		default:

			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}
