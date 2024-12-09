package main

import "fmt"

var a int=5  // global     

// somewhere else is the heap   [heaper]  a


func AFunc() func() {  // sf1
	a := 5
	return func() {
		a++   // escape analysis
		fmt.Println(a)
	}
}

func main()  { // sf0 
	f := AFunc()  // sf1 removed
	f()   // calling the anonymous through "f"
}
