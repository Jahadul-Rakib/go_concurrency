package main

import (
	"log"
	"sync"
)

func main() {
	var number int

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		wg.Done()
		number++
		log.Println(number)
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		number++
		log.Println(number)
	}()
	wg.Wait()
	println("Data value is: ", number)
}
