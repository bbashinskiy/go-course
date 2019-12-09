package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27, //always need a traling come
	} // composite literal
	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	delete(m, "Emprty")
	fmt.Println(m)

	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["Emprty"])

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}
	fmt.Println(m)

}
