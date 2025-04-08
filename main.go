package main

import "fmt"

func main() {
	const USDtoEUR = 0.92
	const USDtoRUB = 85.85
	const EURtoRUB = USDtoEUR * USDtoRUB
}

func getUserInput() string {
	var num = ""
	fmt.Scan(&num)
	return num
}

func getCalculateResult(num int, cash_1 string, cash_2 string) float64 {

}
