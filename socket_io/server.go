package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err :=net.Listen("tcp","localhost:8080")
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		// go use for run different and multiple thread
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for{
		_ , err := io.WriteString(conn, "Print Response from server \n")
		if err != nil {
			log.Println(err.Error())
		}
		time.Sleep(time.Second)
	}
	//handleConnection(conn)
}
