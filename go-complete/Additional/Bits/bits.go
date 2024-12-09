package main

import "fmt"

func clearBit(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

func main() {
	num := 12 //  1100
	fmt.Println("bit three", hasBit(num, 3))

	num = setBit(num, 0) // 1101
	fmt.Println(num)

	num = clearBit(num, 0) // 1100
	fmt.Println(num)

}