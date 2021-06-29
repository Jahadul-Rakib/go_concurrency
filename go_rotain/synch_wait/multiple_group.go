package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var number []int

func addNumbers(num int, wg *sync.WaitGroup) {

	defer wg.Done()

	time.Sleep(time.Duration(1))
	log.Println("Now Process is start: ", num)
	number = append(number, num)

}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go addNumbers(i, &wg)
	}

	wg.Wait()
	fmt.Println(number)
}
