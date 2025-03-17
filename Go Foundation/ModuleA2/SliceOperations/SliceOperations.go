package main

import "fmt"

func main() {
	myslice := []int{1, 2, 3, 4, 5, 6}

	myslice2 := []int{20, 21, 22}

	a := append(myslice, 7, 8, 9)

	b := append(myslice, myslice2...)

	c := append(myslice, myslice2[:2]...)

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)
}
