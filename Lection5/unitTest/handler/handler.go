package handler

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	text := "handler message"
	if _, err := w.Write([]byte(text)); err != nil {
		log.Fatalln(err)
	}
}
