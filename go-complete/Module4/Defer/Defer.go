package main

import "fmt"

func main() {

	defer fileCleanup()
	defer releaseConnections()

	fmt.Println("in main...")
}

func fileCleanup() {
	fmt.Println("doing file cleanup")
}

func releaseConnections() {
	fmt.Println("releasing connection")
}
