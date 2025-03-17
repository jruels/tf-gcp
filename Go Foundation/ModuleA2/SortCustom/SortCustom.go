package main

import (
	"fmt"
	"sort"
)

type byLength []string // cannot change the behavior of an intrinsic type

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i] //  bubble sort
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
