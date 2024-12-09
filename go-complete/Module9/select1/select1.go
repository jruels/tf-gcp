package main

import (
	"fmt"
	"sync"
)

type distance int
type degree int

var wg = sync.WaitGroup{}

func main() {

	walk := make(chan distance)
	run := make(chan distance)
	turn := make(chan degree)

	wg.Add(2)
	go func() {
		defer wg.Done()
		walk <- 5
		walk <- 20
		turn <- 45
		run <- 5
		walk <- 0
	}()
	go func() {
		defer wg.Done()
		for {
			select {
			case walking := <-walk:
				if walking == 0 {
					fmt.Printf("person has stopped\n")
					return
				}
				fmt.Printf("walking %d blocks\n", walking)
			case running := <-run:
				fmt.Printf("running %d blocks\n", running)
			case turning := <-turn:
				fmt.Printf("turning %d degrees\n", turning)
			}
		}
	}()

	wg.Wait()
}