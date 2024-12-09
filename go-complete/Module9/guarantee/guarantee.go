package main

import (
	"fmt"
	"time"
)

func main() {

	//channel := make(chan int) // unbuffered -guaranteed delivery
	channel := make(chan int, 1) // buffered - delayed delivery

	go func() {
		fmt.Println("about to block")
		channel <- 10 // not going to block
		fmt.Println("after 1st insertion")
		channel <- 11 // if channel not empty - block - delayed delivery
		fmt.Println("delayed insertion")

	}()

	time.Sleep(time.Second * 10)
	value := <-channel
	time.Sleep(time.Second * 5)
	fmt.Println(value)
}
