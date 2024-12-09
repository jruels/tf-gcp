package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	channel := make(chan int)         // bidirectional
	var rchannel <-chan int = channel // read only
	var wchannel chan<- int = channel // write only

	go func() {
		defer wg.Done()
		fmt.Println(<-rchannel)
		fmt.Println(<-rchannel)
		// rchannel <- 1  not compile
	}()

	go func() {
		defer wg.Done()
		wchannel <- 4
		wchannel <- 2
	}()

	wg.Wait()
}
