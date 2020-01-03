package main

import "fmt"

<<<<<<< HEAD
// func (r receiver) identifier(parameters) (returns(s)) {...}

func main() {
	foo()
}

func foo() {
	fmt.Println("hello from foo")
=======
func main() {
	foo()
	bar("James Bond")
	s1 := woo("Maneypenny")
	fmt.Println(s1)
	x, y := mause("Ian", "Flaming")
	fmt.Println(x, "\n", y)
}

func foo() {
	fmt.Println("Hello from foo")
}

// everything in GO is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hellow from woo,", st)
}

func mause(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := false
	return a, b
>>>>>>> 417454f733316d059504a69096228a4790f21141
}
