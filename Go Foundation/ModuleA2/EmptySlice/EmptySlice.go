package main

import "fmt"

func main() {

	var fslice []float64
	fslice = append(fslice, 1.0)
	fmt.Println(fslice)

	var mymap map[string]int
	mymap["test"] = 1
}
