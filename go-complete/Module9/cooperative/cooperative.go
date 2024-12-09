package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	current := runtime.GOMAXPROCS(1)
	fmt.Printf("Cores %d\n", current)
	go func() {
		for i := 1; i < 100; i = i + 2 {
			fmt.Printf("%d ", i)
			runtime.Gosched()
		}
	}()
	go func() {
		for i := 2; i < 100; i = i + 2 {
			fmt.Printf("%d ", i)
			runtime.Gosched()
		}
	}()

	time.Sleep(time.Second)
}
