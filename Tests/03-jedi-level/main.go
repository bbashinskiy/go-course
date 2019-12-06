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
	year := 1992
	for year <= 2019 {
		fmt.Println(year)
		year++
	}
	//4
	bd := 1992
	for {
		if bd > 2019 {
			break
		}
		fmt.Println(bd)
		bd++
	}
	//5
	for dg := 10; dg <= 100; dg++ {
		fmt.Println(dg % 4)
	}
	//6

}
