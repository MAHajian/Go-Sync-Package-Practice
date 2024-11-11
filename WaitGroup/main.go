package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	start := time.Now()
	wg := new(sync.WaitGroup)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			task(id)
		}(i)
	}
	wg.Wait()
	fmt.Printf("Elapsed: %v", time.Since(start))
}

func task(id int) {
	fmt.Printf("doing task # %v\n", id)
}
