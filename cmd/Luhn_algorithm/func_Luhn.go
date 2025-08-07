package main

import (
	"fmt"
	"strings"
)

// Проверяет корректность номера карты, если не правильный - возвращает верную последнюю(16-ю) цифру
func Luhn(number string) (is_valid bool, true_last_digit int) {

	// починить ошибку | trim лучше или все таки цикл для проверки?
	if len(number) != 16 || len(strings.Trim(number, "0123456789")) > 0 {
		fmt.Println("error != 16")
		return
	}

	numbers := make([]int, len(number))
	total_sum := 0

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
		total_sum += num
	}

	if total_sum%10 == 0 {
		is_valid = true
	} else {
		//Считаем последнюю цифру
		if (total_sum-numbers[15])%10 != 0 {
			true_last_digit = 10 - (total_sum-numbers[15])%10
		}

	}

	return
}
