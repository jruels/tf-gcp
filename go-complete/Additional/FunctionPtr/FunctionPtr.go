package main

import "fmt"
import "os"

type fptr func(int) int  // fptr is function ptr type

func main() {

	//var a fptr=double
	var b fptr=triple

	a:=double

	b(4)
	a(15)



	fmt.Println(b)

	if len(os.Args) == 1 {
		funcExecute(double)
	} else {
		funcExecute(triple)
	}


	fmt.Println(b)
	fmt.Println(b)
	fmt.Println(b)

}

func double(a int) int {
	return a * a
}

func triple(a int) int {
	return a * a * a
}


func test(a, b int) int {
	return a * a * a
}

func funcExecute(f fptr) {
	fmt.Println(f(5))
}
