package main

func main() {
	const a int = 5

	const b int = 10 * a

	// Does not work
	var d int = 10
	const c int = 10 * d
}
