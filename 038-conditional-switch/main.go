package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("should not print")
	case (2 == 4):
		fmt.Println("should not print 2")
	case (3 == 3):
		fmt.Println("should print")
		fallthrough // do not use it
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
	case (5 == 3):
		fmt.Println("false")
		fallthrough
	case (10 == 3):
		fmt.Println("false 2")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
	default:
		fmt.Println("this is default")
	}

	fmt.Println("\n")
	n := "Bond"

	switch n {
	case "moneypenny", "Bond", "Other":
		fmt.Println("miss money or bond or other")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("this is Q")
	default:
		fmt.Println("this is default")
	}
}
