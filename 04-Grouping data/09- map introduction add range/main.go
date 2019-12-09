package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27, //always need a traling come
	} // composite literal
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])
	v, ok := m["Barnabas"] // , ok idiom
	fmt.Println(v, ok)

	if _, ok := m["Barnabas"]; !ok {
		fmt.Println("Key doesn't exist")
	}

	m["Todd"] = 55 //add element

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{1, 4, 5, 6, 7, 8}

	for i, v := range xi {
		fmt.Println(i, v)
	}

}
