package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	fmt.Println(x())
	fmt.Println(bar()())

}

func foo() string {
	s := "Hello World"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
