package main

import (
	"log"
	"sync"
)

func main() {

	var once sync.Once

	load := func() {
		log.Println("Runs only one time")
	}
	once.Do(load)

}
