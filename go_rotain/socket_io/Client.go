package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println(err.Error())
	}
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		log.Println(message)
	}
}
