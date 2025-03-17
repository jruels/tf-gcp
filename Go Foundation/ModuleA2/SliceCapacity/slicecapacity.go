package main

import "fmt"

func main() {

	s1 := make([]int, 4, 5)
	fmt.Println("Len ", len(s1), "Capacity ", cap(s1))

	s1 = append(s1, 5)
	s1 = append(s1, 5)
	fmt.Println("Len ", len(s1), "Capacity ", cap(s1))

	s1 = append(s1, 5)
	s1 = append(s1, 5)
	s1 = append(s1, 5)
	s1 = append(s1, 5)
	s1 = append(s1, 5)
	s1 = append(s1, 5)
	fmt.Println("Len ", len(s1), "Capacity ", cap(s1))
}
