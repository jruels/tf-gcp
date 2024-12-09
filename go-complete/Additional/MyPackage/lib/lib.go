package lib

import "fmt"

// FuncA comment
func FuncA() {
	fmt.Println("FuncA")
	funcb()
}

func funcb() {
	fmt.Println("funcb")
}
