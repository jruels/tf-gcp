// go get github.com/johncgriffin/overflow

package main

import (
	"fmt"
	"math"

	"github.com/JohnCGriffin/overflow"
)

func main() {
	value, ok := overflow.Add(math.MaxInt64, 1)
	if !ok {
		fmt.Println("overflow occured")
	}
	fmt.Println("value ", value)
}
