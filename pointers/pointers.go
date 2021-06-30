package main

import "log"

func main() {
	point := 10
	p := &point

	log.Println(*p)     //by using * point value
	log.Println(&point) //by using & point memory address
}
