package main

import (
	"fmt"
)

type person struct { // person 구조체 선언
	fName string
	lName string
}

func main() {
	// composite literal - struct literal

	p1 := person{"Todd", "Mclead"}
	p2 := person{"Nina", "Simone"}
	fmt.Println(p1)
	fmt.Println(p2)

	a1 := struct {
		bread string
		name  string
	}{
		"German Sheperd",
		"Shasta",
	}
	fmt.Println(a1)
}
