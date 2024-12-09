package main

import "fmt"

func main() {

	a:=1
	a+=5

	b:="var"
	b+=" another var"

	matrix := [4][4]interface{}{{"a", 4, 5, 0},
		{"d", 0, 3, 0},
		{"m", 4, 8, 0},
		{"s", 1, 4, 0},
	}

	functions := map[string]func(int, int) int{
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
