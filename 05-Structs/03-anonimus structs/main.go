package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Miss",
		last:  "Moneypenny",
		age:   21,
	}

	fmt.Println(p1.first, p1.last, p1.age)
}
