package main

import "fmt"

func main() {
	str := "नमस्ते"

	bytes := []byte(str)
	fmt.Println(bytes)

	str2 := string(bytes)
	fmt.Println(str2)

	runes := []rune(str)
	fmt.Println(runes)

	str3 := string(runes)
	fmt.Println(str3)
}
