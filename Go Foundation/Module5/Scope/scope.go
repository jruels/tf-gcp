package main

import (
    "fmt"
)

var i = 5

func main() {

    i := 4.5

    {
        i := "test"
        fmt.Println(i)
    }
    fmt.Println(i)
}

