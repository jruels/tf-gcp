package main

import "fmt"

type transaction struct {
	name   string
	amount uint
}
type balance uint

var customer = map[string]balance{
	"Bob":   0,
	"Ted":   0,
	"Carol": 0,
	"Alice": 0,
}

var deposit = make(chan transaction)
var withdrawal = make(chan transaction)
var done = make(chan bool)

func main() {

	go func() {
		for {
			select {
			case depositing := <-deposit:
				value, exist := customer[depositing.name]
				if exist {
					customer[depositing.name] = value +
						balance(depositing.amount)
				} else {
					panic("invalid customer.")
				}
			case withdrawing := <-withdrawal:
				value, exist := customer[withdrawing.name]
				if exist {
					temp := value - balance(withdrawing.amount)
					if temp < 0 {
						panic("cannot overdraw account")
						continue
					}
					customer[withdrawing.name] = temp
				} else {
					panic("invalid customer.")
				}
			case <-done:
				fmt.Println("exiting")
				return
			}
		}
	}()

	deposit <- transaction{"xyz", 10}
	deposit <- transaction{"Ted", 40}
	withdrawal <- transaction{"Bob", 5}
	deposit <- transaction{"Alice", 25}
	withdrawal <- transaction{"Alice", 10}
	deposit <- transaction{"Bob", 10}
	done <- true
	fmt.Println(customer)
}
