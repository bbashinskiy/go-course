package main

import "fmt"

func main() {
	s := "B"
	fmt.Println(s)

	bs := []byte(s)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%x\n", n)
}
