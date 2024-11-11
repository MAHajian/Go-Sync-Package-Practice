package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int

func main() {
	start := time.Now()
	mu := new(sync.Mutex)
	for i := 0; i < 5; i++ {
		go func() {
			go updateCounter(mu)
		}()
	}
	<-time.After(time.Millisecond)
	fmt.Printf("Elapsed: %v", time.Since(start))
}

func updateCounter(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	counter++
	fmt.Printf("%v\n", counter)
}
