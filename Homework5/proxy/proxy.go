package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

const proxyAddr string = "localhost:9000"

var (
	counter        int    = 0
	firstInstance  string = "http://localhost:8080"
	secondInstacne string = "http://localhost:8082"
)

func main() {
	http.HandleFunc("/", handleProxy)
	log.Fatalln(http.ListenAndServe("localhost:9000", nil))
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	textBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	text := string(textBytes)
	if counter == 0 {
		resp, err := http.Post(firstInstance, "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			log.Fatalln(err)
		}
		counter++
		textBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		fmt.Println(string(textBytes))
		return
	}
	resp, err := http.Post(secondInstacne, "text/plain", bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatalln(err)
	}
	counter++
	textBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(textBytes))
	counter--
}
