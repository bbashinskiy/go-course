package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives address in memory, what is the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // type pointer to an int

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * gives the value stored at an address
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)
}
