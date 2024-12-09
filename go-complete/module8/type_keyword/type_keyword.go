package main

import "fmt"

type mystruct struct {
	first string
	last  string
	age   int
	// afunc func()int
}

// func coolstuf() int {
// 	return 12
// }

func ( /* receiver */ a mystruct /* this */) coolstuff() int {
	fmt.Println(a.first, " ", a.last)
	return 5
}

func ( /* receiver */ obj *mystruct /* this */) birthday() {
	obj.age++
}
func main() {
	obj := mystruct{"bob", "young", 42}
	obj.coolstuff()
	obj.birthday()
	fmt.Println(obj.age)
}
