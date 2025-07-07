package main

import (
	"sync"

	"github.com/singleton-dp/singleton"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		sing1 := singleton.NewSingleton()
		sing1.Write(2)
		wg.Done()
	}()

	go func() {
		sing2 := singleton.NewSingleton()
		sing2.Write(3)
		wg.Done()
	}()

	wg.Wait()
}