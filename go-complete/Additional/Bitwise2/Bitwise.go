package main

import "fmt"

func main() {

	i := 10     // 1010
	j := i >> 1 // 0101

	k := i | j // 1111

	fmt.Println("j", j, "k", k)

}
