package main

import (
	"log"
	"time"
)

func intoFunction(cnl chan int) {
	data := 1532
	cnl <- data
	log.Println("send data to channel: ", data)
}
func dividedFunction(cnl chan int) {
	data := <-cnl
	log.Println("receive data from channel: ", data)
}

func main() {
	var intoChannel = make(chan int)

	go intoFunction(intoChannel)
	go dividedFunction(intoChannel)

	time.Sleep(time.Second * 30)
}
