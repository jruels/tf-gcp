package main

type IA interface {
	FuncA()
	FuncB()
}

func FuncD(obj IA) {

}

type MyStruct struct {

}

func (a MyStruct) FuncB() {


}

func main() {

	var obj MyStruct
	FuncD(obj)
}
