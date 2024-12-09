package main

import "fmt"

func main() {
	var xyz int = 5
	var pxyz *int

	pxyz = &xyz
	fmt.Println(pxyz, *pxyz)
}
