package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{345, 56, 23, 65}
	x = append(x, y...)
	fmt.Println(x)
	x = append(x, y[0])
	fmt.Println(x)

	x = append(x[:6], x[9:]...) // ... extract all data from data type slice int into int
	fmt.Println(x)
}
