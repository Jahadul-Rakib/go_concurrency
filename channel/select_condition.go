package main

import (
	"log"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3)
		ch1 <- "Channel One"
	}()

	go func() {
		time.Sleep(2)
		ch2 <- "Channel Two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case v := <-ch1:
			log.Println(v)
		case v := <-ch2:
			log.Println(v)
		default:
			log.Println("No match")
		}
	}
}
