package main

import (
	"fmt"
)

func main() {
	// slice
	// composite literal - map literal

	// var idMap map[int]string

	x := map[string]int{
		"Todd":    45,
		"Nina":    25,
		"Patrick": 27,
	}

	for k, v := range x {
		fmt.Println(k, "-", v)
	}

	for k, _ := range x {
		fmt.Println(k, "-", x[k])
	}

	// make()
	y := make(map[string]string)
	y["James"] = "007"
	fmt.Println(y["James"])
}
