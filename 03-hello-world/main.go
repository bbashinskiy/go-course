package main

import (
	"fmt"
)

var z int

func main() {
	x := 43
	fmt.Println(x)
	x = 54
	y := x + 32
	fmt.Println(y)
	textnumber := "I will be old in"
	fmt.Println(textnumber, y, "years.")
	fmt.Println(z)
	foo()
}

func foo() {
	fmt.Println("from another function", z)
}
