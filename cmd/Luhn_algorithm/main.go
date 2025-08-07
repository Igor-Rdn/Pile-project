package main

import (
	"fmt"
)

func main() {

	var card_number string
	fmt.Print("Введите номер карты (16 цифр): ")
	fmt.Scanln(&card_number)

	is_valid, true_last_digit := Luhn(card_number)

	if is_valid {
		fmt.Println("The number is valid")
	} else {
		fmt.Printf("The number is not valid, the last digit must be:%d", true_last_digit)
	}
	fmt.Scanln()
}
