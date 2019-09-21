package pointers

import "fmt"

func main() {
	x := 42
	Foo(x)
	fmt.Println(x)

	w := 42
	Bar(&w)
	fmt.Println(w)

	var ps *string
	s := "String"
	ps = &s
	fmt.Println(s)
	fmt.Println(ps)
	fmt.Println(*ps)
}

func Foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func Bar(y *int) {
	fmt.Println(*y)
	*y = 43
	fmt.Println(*y)
}
