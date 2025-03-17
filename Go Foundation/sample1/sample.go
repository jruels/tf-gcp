package main

import "fmt"
import _ "errors"

// variadic ...
func defer_method() {

	// prototype: func recover() interface{}
	//obj := recover()
	// if obj == nil {
	// 	fmt.Println("nothing to handle")
	// 	return
	// }

	// fmt.Printf("Recover obj %T\n", obj)

	obj2 := obj.(error).Error()
	fmt.Printf("Recover obj %T\n", obj2)

}

func afunc() {
	defer defer_method()

	//panic(5)
	//panic("string")

	//message := errors.New("embedded message")
	//panic(message)
}

func main() {
	afunc()
	fmt.Println("in main")
}
