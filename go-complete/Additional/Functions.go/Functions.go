package main

import "fmt"

func main() {

    funcA(5, "cool")
    funcB(5, true, false, false)
}

func funcA(a int, b string) {
    fmt.Println(a, b)
}

func funcB(a int, b ...bool) {
    fmt.Println(b)
}


