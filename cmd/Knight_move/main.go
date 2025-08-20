package main

import (
	"fmt"
)

type position struct {
	x int // i
	y int // j
}

var availableMoves = []position{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}

// Проверка валидности позиции
func (p position) isValid() bool {
	return p.x >= 0 && p.x < 8 && p.y >= 0 && p.y < 8
}

// Определение доступных клеток для следующего шага
func (curPos position) knightMove() (res []position) {

	for _, move := range availableMoves {

		newPos := position{curPos.x + move.x, curPos.y + move.y}

		if newPos.isValid() {
			res = append(res, newPos)
		}
	}
	return
}

func main() {

	startPos := position{2, 3}
	availablePosition := startPos.knightMove()

	fmt.Println(availablePosition)
	fmt.Println(startPos)
	fmt.Println(availableMoves)

}
