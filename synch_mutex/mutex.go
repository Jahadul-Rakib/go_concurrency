package main

import (
	"log"
	"sync"
)

var balance = 0.0

func main() {

	var wg sync.WaitGroup
	var mutex sync.Mutex //used for lock variable value

	deposit := func(amount float64) {
		mutex.Lock()
		balance = balance + amount
		defer mutex.Unlock()
	}
	withdrew := func(amount float64) {
		mutex.Lock()
		balance = balance - amount
		defer mutex.Unlock()

	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		wg.Done()
		deposit(2.23)
	}

	wg.Add(90)
	for i := 0; i < 90; i++ {
		wg.Done()
		withdrew(1.75)
	}
	wg.Wait()

	log.Println("Balance is: ", balance)

}
