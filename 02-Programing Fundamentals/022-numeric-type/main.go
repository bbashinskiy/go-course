package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64

func main() {
	x = 43
	y = 43.4564576
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
