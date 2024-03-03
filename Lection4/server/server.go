package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

func main() {
	listener, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("server start")
	connection, err := listener.Accept()
	if err != nil {
		log.Fatalln(err)
	}
	
	for {
		line, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("line received ", string(line))
		upperLine := strings.ToUpper(string(line))
		if _, err := connection.Write([]byte(upperLine)); err != nil {
			log.Fatalln(err)
		}
	}
}
