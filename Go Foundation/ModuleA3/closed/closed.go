package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int)
	go func() {
		channel <- 5
		time.Sleep(time.Millisecond * 5)
		channel <- 15
	}()

	go func() {
		fmt.Println(<-channel)
		close(channel)
	}()

	time.Sleep(time.Second * 10)
	fmt.Println(<-channel)

}
