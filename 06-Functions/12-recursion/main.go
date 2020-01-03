package main

import "fmt"

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	n := factorial(4)
	fmt.Println(n)
	n2 := factorialloop(4)
	fmt.Println(n2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//4 * 3 * 2 * 1 * 1

func factorialloop(n int) int {
	fact := 1
	for ; n > 0; n-- {
		fact *= n
	}
	return fact
}
