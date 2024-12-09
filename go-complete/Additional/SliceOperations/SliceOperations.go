package main

import "fmt"

func main() {
	myslice := []int{1, 2, 3, 4, 5, 6}

	a := myslice[4]

	b := myslice[2:4]

	c := myslice[2:]

	d := myslice[:4]

	e := myslice[:]

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)

	fmt.Println(d)

	fmt.Println(e)
}
