package main

import (
	"fmt"
	"strconv"
	"time"
)

func printData(c chan string) {
	fmt.Println(<-c)
}

func countTo(max int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < max; i++ {
			value := strconv.Itoa(i)
			ch <- value + "data"
		}
		close(ch)
	}()
	return ch
}
func countToAnd(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			default:
				ch <- i
			}
		}
		close(ch)
	}()
	return ch, cancel
}
func main() {

	var channel = make(chan string)
	go printData(channel)
	channel <- "rakib"
	fmt.Println("---------------")

	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("---------------")

	for i := range countTo(10) {
		fmt.Println(i)
	}
	fmt.Println("---------------")

	cnl, cancel := countToAnd(10)
	for i := range cnl {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	cancel()
	fmt.Println("---------------")

	time.Sleep(10000)
}
