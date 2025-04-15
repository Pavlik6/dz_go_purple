package main

import (
	"errors"
	"fmt"
)

const USDtoEUR = 0.92
const USDtoRUB = 85.85
const EURtoRUB = USDtoEUR * USDtoRUB

const USD string = "USD"
const EUR string = "EUR"
const RUB string = "RUB"

func main() {
	var cashFirst string
	var cashSecond string
	var cashCount int
	var result float64

	fmt.Println("__Калькулятор обмена валюты__")

	for {
		fmt.Printf("1. Введите исходную валюту: (%s, %s, %s)\n", USD, EUR, RUB)
		cash_1, err := getUserInputCash()
		if err != nil {
			fmt.Println("Не верно введено!!!")
			continue
		}
		cashFirst = cash_1
		break
	}

	for {
		fmt.Println("2. Введите кол-во валюты: ")
		numOfCash, err := getUserInputCount()
		if err != nil {
			fmt.Println("Не верно введено!!!")
			continue
		}
		cashCount = numOfCash
		break
	}

	for {
		fmt.Printf("3. Введите целевую валюту: (%s %s %s)\n", USD, EUR, RUB)
		cash_2, err := getUserInputCash()
		if err != nil {
			fmt.Println("Не верно введено!!!")
			continue
		}
		cashSecond = cash_2
		result = getCalculateResult(cashCount, cashFirst, cashSecond)
		fmt.Printf("Общее кол-во денег равно: %.2f\n", result)
		break
	}
}

func getUserInputCash() (string, error) {
	var cash string
	fmt.Scan(&cash)
	if cash != USD && cash != EUR && cash != RUB {
		return "", errors.New("Неверная валюта")
	}
	return cash, nil
}

func getUserInputCount() (int, error) {
	var cashCount int
	fmt.Scan(&cashCount)
	if cashCount < 0 {
		return 0, errors.New("Количество валюты не может быть отрицательным")
	}
	return cashCount, nil
}

func getCalculateResult(num int, cash_1 string, cash_2 string) float64 {
	var resultCash float64
	switch {
	case cash_1 == "USD" && cash_2 == "EUR":
		resultCash = USDtoEUR * float64(num)
	case cash_1 == "USD" && cash_2 == "RUB":
		resultCash = USDtoRUB * float64(num)
	case cash_1 == "EUR" && cash_2 == "RUB":
		resultCash = EURtoRUB * float64(num)
	case cash_1 == "RUB" && cash_2 == "USD":
		resultCash = 1 / USDtoRUB * float64(num)
	case cash_1 == "RUB" && cash_2 == "EUR":
		resultCash = 1 / EURtoRUB * float64(num)
	case cash_1 == "EUR" && cash_2 == "USD":
		resultCash = 1 / USDtoEUR * float64(num)
	}
	return resultCash
}
