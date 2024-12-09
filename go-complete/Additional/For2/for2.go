package main

import (
	"fmt"
)

func test(i int) {
	fmt.Printf("go %d", i)
}

func main() {
	a := 1
	defer test(a)
	a++
	defer test(a)
	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}

	j := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range j {
		fmt.Println(v)
	}

	l := 10

	for l != 0 {
		fmt.Println(l)
		l--
	}

	fmt.Println("end")

}
