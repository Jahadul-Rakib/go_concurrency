package main

import (
	"log"
	"time"
)

func senderFunction(cnl chan string) {
	data := "1532"
	data2 := "1531"
	cnl <- data
	cnl <- data2
	log.Println("send data to channel: ", data)
	log.Println("send data to channel: ", data2)
}
func receiverFunction(cnl chan string) {

	log.Println("receive data from channel: ", <-cnl)
	log.Println("receive data from channel: ", <-cnl)

}

func main() {
	var intoChannel = make(chan string, 2)

	go senderFunction(intoChannel)
	go receiverFunction(intoChannel)

	time.Sleep(time.Second * 30)
}
