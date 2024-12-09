// Version 1

func add(a any, b any) any {
	return a + b
}

// Version 2

func add(a any, b any) any {
	return a.(int) + b.(int)
}
