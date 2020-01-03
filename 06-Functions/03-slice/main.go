package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 98, 9}
	x := sum(xi...)
	fmt.Println("The total is:", x)

}

func sum(x ...int) int {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("fo item in index position", i, "we are no adding", v, "to the total which is now", sum)
	}
	return sum
}

//func (r receiver) identifier(parameters(s)) (return(s)) { code }
