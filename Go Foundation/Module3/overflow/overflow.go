package main

import "math"
import "fmt"

func main() {
	var value int64 = 0
	value = math.MaxInt64
	value++
	fmt.Println("value =", value)
}
