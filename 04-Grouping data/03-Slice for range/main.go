package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[2])

	for i, v := range x {
		fmt.Printf("Index: %v	Value: %v\n", i, v)
	}
}
