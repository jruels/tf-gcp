func add(a any, b any) any {
	if reflect.TypeOf(a) != reaflect.TypeOf(b) {
		panic("incompatible types")
	}
	switch a.(type) {
	case int:
		return a.(int) + b.(int)
	case string:
		return a.(string) + b.(string)
	default:
		panic("not implemented")
	}
	return 0
}

func main() {

}