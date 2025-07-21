package board

import (
	"container/list"
	"fmt"
)

type Board struct {
	Width  int
	Height int
}

type SnakePoint struct {
	SnakePositionX int
	SnakePositionY int
}

func DrawBoard(brd Board, occupiedMap map[SnakePoint]*list.Element) {
	fmt.Print("\033[H\033[2J") // Clear the console
	for y := range brd.Height {
		for x := range brd.Width {
			if x == 0 || x == brd.Width-1 || y == 0 || y == brd.Height-1 {
				fmt.Print("#")
			} else if _, exists := occupiedMap[SnakePoint{SnakePositionX: x, SnakePositionY: y}]; exists {
				fmt.Print("0")
			} else {
				fmt.Print(" ")
			}
		}
		println()
	}
}
