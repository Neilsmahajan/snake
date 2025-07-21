package board

import (
	"fmt"
)

type BoardDimensions struct {
	Width  int
	Height int
}

type SnakePoint struct {
	SnakePositionX int
	SnakePositionY int
}

func DrawBoard(boardDimensions BoardDimensions, body []SnakePoint) {
	fmt.Print("\033[H\033[2J") // Clear the console
	for y := range boardDimensions.Height {
		for x := range boardDimensions.Width {
			if x == 0 || x == boardDimensions.Width-1 || y == 0 || y == boardDimensions.Height-1 {
				fmt.Print("#")
			} else if x == body[len(body)-1].SnakePositionX && y == body[len(body)-1].SnakePositionY {
				fmt.Print("0")
			} else {
				fmt.Print(" ")
			}
		}
		println()
	}
}
