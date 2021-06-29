package main

import (
	"fmt"
	"time"
)

func printData(c chan string) {
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func main() {
	var channel = make(chan string, 3)
	go printData(channel)
	channel <- "rakib"
	channel <- "rakib1"
	channel <- "rakib2"

	time.Sleep(10000)
}
