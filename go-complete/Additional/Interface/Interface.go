package main

type Vehicle interface {
	Start(a int, b int) (int, float64)
	Turn()
	Stop()
}

type Auto struct {
}

func (a Auto) Start() {

}
func (a Auto) Turn() {

}

func (a Auto) Stop() {

}

type Train struct {
}

func (a Train) Start() {

}

func (a Train) Stop() {

}

func turnLeft(v Vehicle) {

}

func main() {
	var a Auto
	var b Train

	turnLeft(a)

	turnLeft(b)

}
