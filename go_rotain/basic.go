package main

import (
	"log"
	"strconv"
	"time"
)

func takeWater(amount int, cnl chan string) <-chan string {
	cnl <- "take " + strconv.Itoa(amount) + " litter."
	return cnl
}

func takeSweetAndMilk(milk int, sweet int, ch chan string) <-chan string {
	ch <- "take milk " + strconv.Itoa(milk) + " and take sweet " + strconv.Itoa(sweet)
	return ch
}

func makingCoffee(waterChannel chan string, milkAndSweetChannel chan string, done chan string) <-chan string {
	log.Println(<-waterChannel)
	log.Println(<-milkAndSweetChannel)
	log.Println("Mixing All of them together.")
	done <- "done"
	return done
}

func main() {
	cnl := make(chan string)
	go takeWater(12, cnl)

	cnl2 := make(chan string)
	go takeSweetAndMilk(6, 3, cnl2)

	doneSignal := make(chan string)
	go makingCoffee(cnl, cnl2, doneSignal)
	log.Println("Congratulations Your Coffee is ready!!!!!! ", <-doneSignal)

	time.Sleep(10000)
}
