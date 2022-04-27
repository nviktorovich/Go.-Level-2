package main

import (
	"fmt"
	"sync"
)

var index int

func main() {
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mx.Lock()
			index++
			mx.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("final index is", index)
}
