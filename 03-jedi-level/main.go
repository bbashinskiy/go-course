package main

import "fmt"

func main() {
	// 1
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
	// 2
	for x := 65; x <= 90; x++ {
		fmt.Println(x)
		for y := 0; y <= 3; y++ {
			fmt.Printf("%#U\n", x)
		}
	}
	//3

}
