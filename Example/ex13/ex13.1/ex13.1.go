package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {

	var house House

	house.Address = "경기도 용인시 죽전동"
	house.Size = 10
	house.Price = 55.5
	house.Type = "오피스텔"

	var house2 House = House{
		Size: 20,
		Type: "주택",
	}

	fmt.Println("주소 : ", house.Address)
	fmt.Println("house2의 타입 : ", house2.Type)
}
