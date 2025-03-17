package main

import "fmt"

func main() {
	var a = 12.55
	var b = int(a)
	var c = float64(b)
	var d = []byte("test")
	var e = []rune("test")

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
