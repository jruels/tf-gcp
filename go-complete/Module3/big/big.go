package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	bigval := new(big.Int) // 1

	bigval.SetInt64(123)

	fmt.Println("bigval = ", bigval)

	op1 := big.NewInt(math.MaxInt64) // 2
	op2 := big.NewInt(math.MaxInt64)
	op3 := bigval.Mul(op1, op2)
	//op3 = bigval.Mul(op3, op2)
	fmt.Println(bigval)
	bigval.Mul(op3, op2)
	fmt.Println(bigval, op1, op2, op3)
}
