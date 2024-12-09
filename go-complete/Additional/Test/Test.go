type fptr func(int, int) interface{}

func add(a int, b int) int { return a + b }
func sub(a int, b int) int { return a - b }
func mul(a int, b int) int { return a * b }
func div(a int, b int) int { return a / b }
func mod(a float64, b float64) float64 { return math.Mod(a, b) }
func abs(a float64) float64 { return math.Abs(a) }
func compute(operation string, a int, b int) int{
    table := map[string]fptr{
        "+": add,
        "/": div,
        "*": mul,
        "-": sub,
        "%": mod,
        "|": abs,
    }
    function, _ := table[operation]
    return function(a, b)
}