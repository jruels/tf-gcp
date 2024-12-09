package main

type Count int
type StringInt map[string]int
type Fptr func(int) (int, int)
type MyStruct struct {
	data1 int
	data2 int
}

func main() {

	var a Count
	var b StringInt
	var c MyStruct
}
