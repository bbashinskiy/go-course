package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	//1
	p1 := person{
		first:   "John",
		last:    "kochiviak",
		flavors: []string{"fruit", "vanilla"},
	}

	p2 := person{
		first:   "Bohdan",
		last:    "Bondarchyk",
		flavors: []string{"fruit", "rom and cola"},
	}

	for _, flavor := range p1.flavors {
		fmt.Println(flavor)
	}
	//2
	mp := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(mp)
	for _, v := range mp {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for _, flavor := range v.flavors {
			fmt.Println(flavor)
		}
	}
	//3
	t1 := truck{
		fourWheel: true,
		vehicle: vehicle{
			doors: 4,
			color: "Red",
		},
	}

	s1 := sedan{
		luxury: true,
		vehicle: vehicle{
			doors: 4,
			color: "Red",
		},
	}

	fmt.Println(s1)
	fmt.Println(s1.luxury)
	fmt.Println(s1.color)
	fmt.Println(s1.doors)
	fmt.Println(t1)
	fmt.Println(t1.fourWheel)
	fmt.Println(t1.color)
	fmt.Println(t1.doors)
	//4
	myself := struct {
		first string
		last  string
		age   int
	}{
		first: "Bohdan",
		last:  "Bondarchyk",
		age:   27,
	}

	fmt.Println(myself)
}
