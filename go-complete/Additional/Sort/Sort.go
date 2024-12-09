package main

import "fmt"
import "sort"

func main() {
	myslice := []int{11, 2, 31, 4, 15, 6}

	sort.Ints(myslice)

	fmt.Println(myslice)

	fmt.Println(sort.IntsAreSorted(myslice))
}
