package main

import "fmt"

func main() {
	r1 := rectangle{10, 20, 40, 30}
	r1.Print()
	r1.Draw()
}

type rectangle struct {
	top    int
	left   int
	bottom int
	right  int
	// func Draw()  --- no in Go
}

func (r /* this */ rectangle) Draw() {  // receiver or target
	fmt.Printf("Top: %d Left: %d Bottom: %d Right: %d",
		r.top, r.left, r.bottom, r.right)
}

func (r rectangle) Print() {
	fmt.Println(r)
}
