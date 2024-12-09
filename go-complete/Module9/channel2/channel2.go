package main

import "fmt"

func main() {

	channel := make(chan int)   // unbuffered channel - blocking

	go func() {

		channel <- 10

	}()

	value := <-channel
	fmt.Println(value)
}
