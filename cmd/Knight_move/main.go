package main

import (
	"fmt"
	"strings"
)

type position struct {
	x int //a -> h
	y int //1 -> 8
}

/* type testCases struct {
	start position
	end   position
} */

// Проверка валидности позиции
func (p position) isValid() bool {
	return p.x >= 0 && p.x < 8 && p.y >= 0 && p.y < 8
}

// Определение доступных клеток для следующего шага
func (curPos position) knightMove() (res []position) {

	var availableMoves = []position{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}

	for _, move := range availableMoves {

		newPos := position{curPos.x + move.x, curPos.y + move.y}

		if newPos.isValid() {
			res = append(res, newPos)
		}
	}
	return
}

// Поиск кратчайшего пути
func findShortestPath(startPos, endPos position) []position {

	queue := []position{startPos}
	parent := make(map[position]position) //[Child]Parent

	for len(queue) > 0 {

		curPos := queue[0]

		if curPos == endPos {
			break
		}

		for _, availablePos := range curPos.knightMove() {
			if _, exists := parent[availablePos]; !exists && availablePos != startPos {
				queue = append(queue, availablePos)
				parent[availablePos] = curPos
			}
		}
		queue = queue[1:]
	}
	return recoveryPath(endPos, parent)
}

// Восстановление итогового пути
func recoveryPath(parentPos position, parent map[position]position) (path []position) {

	path = append(path, parentPos)

	for {

		if childPos, exists := parent[parentPos]; exists {
			path = append([]position{childPos}, path...)
			parentPos = childPos
		} else {
			break
		}
	}
	return
}

// Шахматные координаты в числовые
func chessToNumeric(notation string) (position, error) {

	if len(notation) != 2 {
		return position{}, fmt.Errorf("неверные координаты: %s", notation)
	}

	lowerNotation := strings.ToLower(notation)
	x := int(lowerNotation[0] - 'a')
	y := int(lowerNotation[1] - '1')

	res := position{x, y}

	if !res.isValid() {
		return position{}, fmt.Errorf("неверные координаты: %s", notation)
	}

	return res, nil
}

// Числовые координаты в шахматные
func numericToChess(pos position) string {
	return fmt.Sprintf("%c%d", 'a'+pos.x, pos.y+1)
}

// Вывод пути
func formatPath(path []position) []string {
	var result []string
	for _, pos := range path {
		result = append(result, numericToChess(pos))
	}
	return result
}

func main() {

	var startPosString string
	var endPosString string

	fmt.Println("Введите начальные координаты:")
	fmt.Scanln(&startPosString)
	fmt.Println("Введите конечные координаты:")
	fmt.Scanln(&endPosString)

	startPos, err := chessToNumeric(startPosString)
	if err != nil {
		fmt.Printf("Ошибка! %v", err)
		fmt.Scanln()
		return
	}

	endPos, err := chessToNumeric(endPosString)
	if err != nil {
		fmt.Printf("Ошибка! %v", err)
		fmt.Scanln()
		return
	}

	shortestPath := findShortestPath(startPos, endPos)
	fmt.Println("Количество шагов:", len(shortestPath)-1)
	fmt.Println("Путь:", formatPath(shortestPath))

	fmt.Scanln()

	/* 	testCases := []testCases{
		{position{3, 5}, position{6, 1}},
		{position{0, 0}, position{7, 7}},
		{position{0, 0}, position{2, 1}},
		{position{0, 0}, position{0, 0}},
	} */

	/* 	for _, tc := range testCases {
		shortestPath := findShortestPath(tc.start, tc.end)
		fmt.Printf("Тест %d: %v -> %v\n", i+1, tc.start, tc.end)
		fmt.Printf("Количество ходов %d: %v\n\n", len(shortestPath)-1, shortestPath)

	} */

}
