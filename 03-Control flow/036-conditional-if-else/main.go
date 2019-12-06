package main

import (
	"fmt"
)

func main() {
	x := 42
	if x == 40 {
		fmt.Println("our value was", x)
	} else if x == 41 {
		fmt.Println("our value was", x)
	} else {
		fmt.Println(x)
	}
}
