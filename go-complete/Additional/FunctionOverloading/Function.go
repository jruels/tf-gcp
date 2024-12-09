package main

import (
	"fmt"
)

func FuncA() {

}

func FuncA(a int) {
	fmt.Println(a)
}

func main() {

	FuncA(5)
}
