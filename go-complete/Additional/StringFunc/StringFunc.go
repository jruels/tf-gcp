package main

import "fmt"

func main() {
	fmt.Println(stringModify("Test", addExclamation))
	fmt.Println(stringModify("Test", right1))
}

type modifier func(string) string

func stringModify(s string, f1 modifier) string {
	return f1(s)
}

func addExclamation(s string) string {
	return s + "!"
}

func right1(s string) string {
	var temp string
	for _, char := range s { // byte array | only ANSI
		temp = temp + string(char+1)
	}
	return temp
}
