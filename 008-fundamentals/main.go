package main

import "fmt"

const (
	one      = 43
	two bool = true
)

const (
	_              = iota
	one_year_ago   = 2019 - iota
	two_year_ago   = 2019 - iota
	three_year_ago = 2019 - iota
	foor_year_ago  = 2019 - iota
)

func main() {
	x := 38
	fmt.Printf("%d\t\t%b\t\t%#x\n\n", x, x, x)

	y := 26
	a := (x == y)
	b := (x <= y)
	c := (x >= y)
	d := (x != y)
	e := (x > y)
	f := (x < y)
	fmt.Printf("%T\t%v\n", a, a)
	fmt.Printf("%T\t%v\n", b, b)
	fmt.Printf("%T\t%v\n", c, c)
	fmt.Printf("%T\t%v\n", d, d)
	fmt.Printf("%T\t%v\n", e, e)
	fmt.Printf("%T\t%v\n", f, f)

	fmt.Println(one, two)

	zz := 34
	fmt.Printf("%d\t\t%b\t\t%#x\n", zz, zz, zz)
	zzl := zz << 1
	fmt.Printf("%d\t\t%b\t\t%#x\n", zzl, zzl, zzl)

	s := `
		One
		ytwo theree 
	test`
	fmt.Println(s)

	fmt.Println(one_year_ago)
	fmt.Println(two_year_ago)
	fmt.Println(three_year_ago)
	fmt.Println(foor_year_ago)
}
