package main

import "fmt"

type fptr func(int, int) int

func main() {

	matrix := [4][4]any{
		{"d", 0, 3, 0},
		{"m", 4, 8, 0},
		{"s", 1, 4, 0},
		{"a", 2, 4, 0},
	}

	functions := map[string]fptr{
		"a": func(a int, b int) int { return a + b },
		"d": func(a int, b int) int { return a / b },
		"m": func(a int, b int) int { return a * b },
		"s": func(a int, b int) int { return a - b },
	}

	for index, row := range matrix {
		matrix[index][3] = functions[((row[0]).(string))](row[1].(int), row[2].(int))
	}

	fmt.Println(matrix)
}
