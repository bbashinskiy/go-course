package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("\n%T\t%b\t%x", y, y, y)
	s := fmt.Sprintf("\n%T\t%b\t%x", y, y, y)
	fmt.Printf(s)
	fmt.Printf("%v", s)
}
