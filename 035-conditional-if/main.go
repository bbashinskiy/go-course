package main

import "fmt"

func main() {

	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
	if 2 == 2 {
		fmt.Println("005")
	}
	if 2 != 2 {
		fmt.Println("006")
	}
	if !(2 == 2) {
		fmt.Println("007")
	}
	if !(2 != 2) {
		fmt.Println("008")
	}

	if x := 42; x == 42 {
		fmt.Println("009")
		fmt.Println(x)
	}
	fmt.Println("here's a statment")
	fmt.Println("something else")
	// doesn't work x scope limited
	// fmt.Println(x)
}
