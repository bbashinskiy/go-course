package main

import "fmt"

func main() {
	// 1
	x := [5]int{23, 436, 567, 567, 567}

	for i, v := range x {
		fmt.Printf("Index: %v, Type: %T, Value: %v\n", i, v, v)
	}
	//2
	y := []int{23, 436, 567, 567, 567, 45, 56, 868, 98, 89, 676}

	for i, v := range y {
		fmt.Printf("Index: %v, Type: %T, Value: %v\n", i, v, v)
	}
	//3
	z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(z)
	fmt.Println(z[:5])
	fmt.Println(z[5:])
	fmt.Println(z[2:7])
	fmt.Println(z[1:6])
	//4
	z = append(z, 52)
	fmt.Println(z)
	z = append(z, 53, 54, 55)
	fmt.Println(z)
	m := []int{56, 57, 58, 59, 60}
	z = append(z, m...)
	fmt.Println(z)
	//5
	k := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	k = append(k[:3], k[6:]...)
	fmt.Println(k)
	//6
	states := make([]string, 50, 50)
	states = []string{`Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	fmt.Println(len(states))
	fmt.Println(cap(states))
	fmt.Println(states)

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}
	//7
	fr := []string{"James", "Bond", "Shaken, not stirred"}
	sr := []string{"Miss", "Moneypenny", "Hello, James"}
	dm := [][]string{fr, sr}
	fmt.Println(dm)
	for _, r := range dm {
		for _, v := range r {
			fmt.Printf(v)
		}
	}
	//8
	mp := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println("\n")
	for k, v := range mp {
		fmt.Println("Key:", k)
		for i, el := range v {
			fmt.Printf("\tIndex: %v,\tValue: %v\n", i, el)
		}
	}
	//9
	fmt.Println("\n")
	mp["New record"] = []string{"test", "record"}
	for k, v := range mp {
		fmt.Println("Key:", k)
		for i, el := range v {
			fmt.Printf("\tIndex: %v,\tValue: %v\n", i, el)
		}
	}
	//10
	fmt.Println("\n")
	delete(mp, "no_dr")

	for k, v := range mp {
		fmt.Println("Key:", k)
		for i, el := range v {
			fmt.Printf("\tIndex: %v,\tValue: %v\n", i, el)
		}
	}
}
