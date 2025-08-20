package main

import (
	"fmt"
)

type position struct {
	x int // i
	y int // j
}

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

func findShortestPath(startPos, endPos position) []position {

	queue := []position{startPos}
	parent := make(map[position]position) //[Child]Parent
	//i := 0

	for len(queue) > 0 {
		//i++

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
		//fmt.Println("i", i, queue)

	}
	return reconstructPath(endPos, parent)
}

func reconstructPath(parentPos position, parent map[position]position) (resSlice []position) {

	resSlice = append(resSlice, parentPos)

	for {
		if childPos, exists := parent[parentPos]; exists {
			resSlice = append([]position{childPos}, resSlice...)
			parentPos = childPos
		} else {
			break
		}
	}
	return
}

func main() {

	startPos := position{2, 3}
	endPos := position{6, 6}

	shortestPath := findShortestPath(startPos, endPos)
	fmt.Println(shortestPath)
}
