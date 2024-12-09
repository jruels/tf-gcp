package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() { // thread
		time.Sleep(2 * time.Second)
		channel1 <- "go routine1"
	}() // thread is invoked
	go func() { // thread
		time.Sleep(4 * time.Second)
		channel2 <- "go routine2"
	}() // thread is invoked

	for i := 0; i < 2; i++ {
		select {
		case message1 := <-channel1:
			fmt.Println("Channel1", message1)
		case message2 := <-channel2:
			fmt.Println("Channel2", message2)
		}
	}
}
