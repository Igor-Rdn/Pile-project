package main

import (
	"fmt"
)

func main() {

	var cardNumber string
	fmt.Print("Введите номер карты (16 цифр): ")
	fmt.Scanln(&cardNumber)

	isValid, trueLastDigit := Luhn(cardNumber)

	if isValid {
		fmt.Println("The number is valid")
	} else {
		fmt.Printf("The number is not valid, the last digit must be:%d", trueLastDigit)
	}
	fmt.Scanln()
}
