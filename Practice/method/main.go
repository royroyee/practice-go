package main

import "fmt"

type person struct {
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func main() {

	p1 := person{
		"Miss",
		"Moneypenny",
	}

	// fmt.Println(p1)
	p1.speak() // use method
}
