package main

import "fmt"

type IPerson interface {
	talk()
	walk()
	eat()
}

type Human struct {
}

type Martian struct {
}

func (Human) talk() {
	fmt.Println("talking")
}

func (Human) walk() {
	fmt.Println("walking")
}

func (Human) eat() {
	fmt.Println("eating sushi")
}

func (Martian) talk() {
	fmt.Println("talking")
}

func (Martian) walk() {
	fmt.Println("walking")
}

func (Martian) eat() {
	fmt.Println("eating Humans")
}

func invite2Party(partygoers ...IPerson) {

	for _, person := range partygoers {
		person.eat() // duck typing   ------   unrelated
	}
}

func main() {
	invite2Party(Martian{}, Human{}, Martian{}, Human{})
}
