package main

import (
	"log"
	"sync"
)

type SlowComplicatedParser struct {
	name string
}

var parser SlowComplicatedParser
var once sync.Once

func Parse() SlowComplicatedParser {
	once.Do(func() {
		parser = initParser()
		log.Println("One")
	})
	return parser
}

func initParser() SlowComplicatedParser {
	return SlowComplicatedParser{name: "Rakib"}
}

func main() {
	defer Parse()

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		wg.Done()
		parse := Parse()
		log.Println("Called ", parse.name)
	}()
	go func() {
		wg.Done()
		complicatedParser := Parse()
		log.Println("Called", complicatedParser.name)
	}()

	wg.Wait()
}
