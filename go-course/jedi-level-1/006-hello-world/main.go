package main

import "fmt"

var x int

type myown int

var y myown

func main() {
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y = 54
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	x = int(y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
