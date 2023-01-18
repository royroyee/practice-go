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

type human interface {
	speak() // person과 secretagent 는 모두 speack를 묵시적으로 구현하고 있다
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		"Miss",
		"Moneypenny",
	}

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	saySomething(p1)
	saySomething(sa1)

}
