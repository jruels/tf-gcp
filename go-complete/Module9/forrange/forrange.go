package main

import "fmt"
import "time"

func main() {
	queue := make(chan bool, 5)
	go func() {
		queue <- true
		queue <- true
		queue <- false
		queue <- true
		queue <- false

		//close(queue)
	}()

	go func() {
		for item := range queue {
			fmt.Println(item)
		}
		fmt.Println("after")
	}()

	time.Sleep(time.Second * 2)
}
