package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	res := calculateAverage(input)

	fmt.Println("res", res)

}

func calculateAverage(input string) string {

	count := 0
	total := 0

	for val := range strings.SplitSeq(input, " ") {

		grade, err := strconv.Atoi(val)
		if err != nil || grade < 0 || grade > 100 {
			continue
		}

		count++
		total += grade

	}

	if count == 0 {
		return "Нет корректных оценок"
	}

	res := float64(total) / float64(count)

	return fmt.Sprintf("%.1f", res)
}
