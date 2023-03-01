package main

import "fmt"

func PrintAvgScore(name string, math int, eng int, history int) {
	var total int = math + eng + history
	avg := total / 3
	fmt.Println(name, "님의 평균 점수 : ", avg, "입니다.")
}

func main() {

	PrintAvgScore("younghwan", 100, 100, 100)

}
