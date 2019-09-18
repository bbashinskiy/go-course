package main

import "fmt"

var y = 42
var x = "Test"
var z = `I said, 

"One two"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
