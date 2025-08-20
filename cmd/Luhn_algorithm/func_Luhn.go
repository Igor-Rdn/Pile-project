package main

import (
	"fmt"
	"strings"
)

// Проверяет корректность номера карты, если не правильный - возвращает верную последнюю(16-ю) цифру
func Luhn(number string) (isValid bool, trueLastDigit int) {

	// починить ошибку | trim лучше или все таки цикл для проверки?
	if len(number) != 16 || len(strings.Trim(number, "0123456789")) > 0 {
		fmt.Println("error != 16")
		return
	}

	numbers := make([]int, len(number))
	totalSum := 0

	// Заполняем слайс цифрами
	for idx, num := range number {
		numbers[idx] = int(num - '0') //numbers[idx], _ = strconv.Atoi(string(num)) ?
	}

	//Алгоритм
	for idx, num := range numbers {
		if idx%2 == 0 {
			if num*2 > 9 {
				num = num*2 - 9
			} else {
				num = num * 2
			}
		}
		totalSum += num
	}

	if totalSum%10 == 0 {
		isValid = true
	} else {
		//Считаем последнюю цифру
		if (totalSum-numbers[15])%10 != 0 {
			trueLastDigit = 10 - (totalSum-numbers[15])%10
		}

	}

	return
}
