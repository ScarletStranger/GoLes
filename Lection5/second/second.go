package main

import (
	"io"
	"log"
	"net/http"
)

const addr string = "localhost:8082"

func main() {
	http.HandleFunc("/", handle)
	log.Fatalln(http.ListenAndServe(addr, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(bodyBytes)
	response := "2 instace: " + text + "\n"

	if _, err := w.Write([]byte(response)); err != nil {
		log.Fatalln(err)
	}
}
