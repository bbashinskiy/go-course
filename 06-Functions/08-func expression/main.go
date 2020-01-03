package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Anonimys ran")
	}
	f()

	g := func(x int) {
		fmt.Println("The year", x)
	}
	g(1942)

}
