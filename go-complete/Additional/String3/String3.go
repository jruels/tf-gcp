package main

import "fmt"

func main() {

	fmt.Println("test1\t1\t2\ntest2\t3\t4\n")
	fmt.Println(`test1\t1\t2\ntest2\t3\t4
        .....
        cool!`)
	fmt.Println("\u3041")
	fmt.Println("\x78")
}
