package main

import "fmt"
import "time"

func main() {

	channel := make(chan int)

	go func() {

		channel <- 10
		fmt.Println("after insertion")

	}()

	time.Sleep(time.Second * 20)
	value := <-channel
	fmt.Println(value)
}
