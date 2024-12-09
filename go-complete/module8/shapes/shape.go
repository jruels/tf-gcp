package main

import "fmt"

type iShape interface { // base type
	draw()
}

type xyz interface { // base type
	draw()
}

type rectangle struct { // derived type
}

func (rectangle) draw() {
	fmt.Println("Drawing rectangle", 5, 6.7)
}

type circle struct { // derived type
}

func (circle) draw() {
	fmt.Println("Drawing circle")
}

func draw_shapes(shapes ...interface{}) { // variadic - slice
	for _, shape := range shapes {
		shape.draw() //  polymorphic - type cast
	}
}

func main() {
	draw_shapes(
		circle{},
		rectangle{},
		circle{},
		rectangle{},
		circle{},
	)
}
