package main

import "fmt"

var x [5]int
var y [6]int //different type

func main() {
	fmt.Println(x)
	x[3] = 43
	fmt.Println(x)
	fmt.Println(len(x))
}
