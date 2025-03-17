package main

import "fmt"
import "math"
import "math/big"

func main() {
	bigval := new(big.Int)

	bigval.SetInt64(123)

	fmt.Println("bigval = ", bigval)

	op1 := big.NewInt(math.MaxInt64)
	op2 := big.NewInt(math.MaxInt64)
	op3 := bigval.Mul(op1, op2)
	op3 = bigval.Mul(op3, op2)
	op3 = bigval.Mul(op3, op2)
	fmt.Println(bigval, op1, op2, op3)
}
