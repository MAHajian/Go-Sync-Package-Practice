package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type singleton struct {
	number int
}

var instance *singleton

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	once := new(sync.Once)
	s1 := createInstance(once)
	s2 := createInstance(once)
	fmt.Printf("%v", s1.number == s2.number)
}

func createInstance(once *sync.Once) *singleton {
	once.Do(func() {
		instance = &singleton{rand.Intn(100)}
	})
	return instance
}
