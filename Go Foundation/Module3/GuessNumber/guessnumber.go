package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	dt := time.Now()

	seed := rand.NewSource(int64(dt.Second()))
	generator := rand.New(seed)
	random := generator.Intn(5) + 1 // inclusive : exclusive

	fmt.Print("Guess a number between 1 and 5: ")

	var reply string
	fmt.Scanln(&reply)
	value, err := strconv.Atoi(reply)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if value == random {
		fmt.Println("You win!!")
	} else {
		fmt.Println("You guess wrong number")
		fmt.Println("The number was", random)
	}

}
