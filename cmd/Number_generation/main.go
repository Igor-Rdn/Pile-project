package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func convertNotation(numeric, notation int, symbols string, count_rank int, zero_symbol string, mask_const string) (res string) {

	for numeric > 0 {

		res = string(symbols[numeric%notation]) + res

		numeric = numeric / notation
	}

	if count_rank > len(res) {
		for range count_rank - len(res) {
			res = zero_symbol + res
		}
	}

	res = mask_const + res

	return res
}

func main() {

	fmt.Println(time.Now())

	rangeFrom := 0
	rangeTo := 2176782335

	notation := 36
	count_rank := 6
	count_required_numbers := 10000000
	symbols := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ" //"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	zero_symbol := "0"
	mask_const := "rigo"

	//var randomNum []int = []int{10, 9838, 12114, 34736, 274298, 522108, 103000000}

	randomNum := make([]int, count_required_numbers) //попробовать неопределенного размера слайс через аппенд
	for i := range count_required_numbers {
		randomNum[i] = randomInt(rangeFrom, rangeTo)
	}

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла")
		return
	}
	defer file.Close()

	sort.Ints(randomNum)
	writer := bufio.NewWriter(file)
	for _, val := range randomNum {
		fmt.Fprintf(writer, "%d:%s\n", val, convertNotation(val, notation, symbols, count_rank, zero_symbol, mask_const))
	}

	writer.Flush()

	fmt.Println(time.Now())
}

/*
		10 = rigo00000A
		9838 = rigo0007LA
		12114 = rigo0009CI
		34736 = rigo000QSW
		274298 = rigo005VNE
		522108 = rigo00B6V0
		103000000 = rigo1PBNB4

		0123456789АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ
		0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ

	}
*/
