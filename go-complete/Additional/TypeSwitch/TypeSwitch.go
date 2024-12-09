package main

import "fmt"

type MyStruct struct {
	a int
	b int
}

func main() {
	var xyz MyStruct

	var val interface{} = xyz

	fmt.Printf("%T\n", val)

//	val.a = 5 // error!!!
	switch val.(type) {
	case string:
		fmt.Println("string")
	case int, int16, int32, int64:
		fmt.Println("int")
	case MyStruct:
		fmt.Println("MyStruct")
	}

	fmt.Println(xyz)
}
