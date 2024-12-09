package main

import "sync"

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go func() {
		wg.Done()
	}()
	go func() {
		wg.Done()
	}()

	wg.Wait()
}
