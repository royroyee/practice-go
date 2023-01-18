package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says. "Shakenn, not stirred."`)
}

func main() {

	p1 := person{
		"Miss",
		"Moneypenny",
	}

	// fmt.Println(p1)
	p1.speak() // use method

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	sa1.speak()        // secret agent speack 호출
	sa1.person.speak() // person speak 호출
}
