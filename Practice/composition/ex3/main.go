package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) pSpeak() {
	fmt.Println(p.fname, `says, "Wassup!"`)
}

func (s secretAgent) saSpeak() {
	fmt.Println(s.fname, "has a license to kill:", s.ltk)
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("uptown, func you up, uptown func you up")
}

func (sa secretAgent) speak() {
	fmt.Println("shaken, not stirred")
}

func speak(h human) {
	h.speak()
}

func main() {
	p1 := person{"Todd", "Mclead"}
	s1 := secretAgent{person{"James", "bond"}, true}
	fmt.Println(p1.fname)
	p1.pSpeak()
	fmt.Println(s1.fname)
	s1.saSpeak()
	s1.pSpeak()

	fmt.Println("Use Interface function")
	p1.speak()
	s1.speak()
}
