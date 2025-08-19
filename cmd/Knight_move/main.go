package main

import (
	"fmt"
)

type position struct {
	x int // i
	y int // j
}

var available_moves = []position{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}

// Проверка валидности позиции
func (p position) is_valid() bool {
	return p.x >= 0 && p.x < 8 && p.y >= 0 && p.y < 8
}

// Определение доступных клеток для следующего шага
func (cur_pos position) knight_move() (res []position) {

	for _, move := range available_moves {

		new_pos := position{cur_pos.x + move.x, cur_pos.y + move.y}

		if new_pos.is_valid() {
			res = append(res, new_pos)
		}
	}
	return
}

func main() {

	start_pos := position{2, 3}
	available_position := start_pos.knight_move()

	fmt.Println(available_position)
	fmt.Println(start_pos)
	fmt.Println(available_moves)

}
