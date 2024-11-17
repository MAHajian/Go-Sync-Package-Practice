package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sharedResourse = map[string]interface{}{}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	start := time.Now()
	cond := sync.NewCond(new(sync.Mutex))
	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		defer wg.Done()
		fetchResource(cond, "res0")
	}()

	go func() {
		defer wg.Done()
		fetchResource(cond, "res1")
	}()

	<-time.After(time.Second * 3)
	
	cond.L.Lock()
	sharedResourse["res0"] = rand.Intn(100)
	sharedResourse["res1"] = rand.Intn(100)
	cond.Broadcast()
	cond.L.Unlock()

	wg.Wait()

	fmt.Println("Elapsed:", time.Since(start))
}

func fetchResource(cond *sync.Cond, key string) {
	defer cond.L.Unlock()
	cond.L.Lock()
	for len(sharedResourse) == 0 {
		cond.Wait()
	}
	fmt.Println(key+":", sharedResourse[key])
}
