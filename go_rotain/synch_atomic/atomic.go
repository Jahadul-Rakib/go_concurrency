package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			atomic.AddInt64(&counter, int64(i))
		}()
	}
	wg.Wait()
	log.Println(counter)
}
