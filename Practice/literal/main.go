package main

import "fmt"

type person struct {
	fname string
	lname string
}

func main() {
	xi := []int{2, 6, 4, 3, 12} // slice
	fmt.Println(xi)

	m := map[string]int{
		"Todd":     45,
		"royroyee": 24,
	}
	fmt.Println(m)

	p1 := person{
		"Miss",
		"Moneypenny",
	}

	fmt.Println(p1)
}
