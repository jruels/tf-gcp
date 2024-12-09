package main

import "fmt"

func AFunc() bool {
	return true
}

func BFunc() bool {
	return false
}

func main() {

	var bool1 = AFunc() || BFunc()
	var bool2 = AFunc() || true
	var bool3 = false && BFunc()
	var bool4 = AFunc()
	var bool5 = BFunc() || AFunc() || false

	fmt.Println(bool1, bool2, bool3, bool4, bool5)
}
