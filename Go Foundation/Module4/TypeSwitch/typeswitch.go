package main

import "fmt"

type MyStruct struct {
	a int
	b int
}

func main() {
	var xyz MyStruct

	var val any = xyz // know very little about val
	// val.a = 5 // error!!!
	switch val.(type) { // RTTI
	case string:
		fmt.Println("string")
	case int, int16, int32, int64:
		fmt.Println("int")
	case MyStruct:
		fmt.Println("MyStruct")
	}

	fmt.Println(xyz)
}
