package board

import (
	"container/list"
	"fmt"

	"github.com/neilsmahajan/snake/internal/types"
)

func DrawBoard(brd *types.Board, occupiedMap map[types.Point]*list.Element) {
	fmt.Print("\033[H\033[2J") // Clear the console
	for y := range brd.Height {
		for x := range brd.Width {
			if x == 0 || x == brd.Width-1 || y == 0 || y == brd.Height-1 {
				fmt.Print("#")
			} else if _, exists := occupiedMap[types.Point{X: x, Y: y}]; exists {
				// Check to see if snake eats fruit
				if _, exists := brd.Fruits[types.Point{X: x, Y: y}]; exists {
					// ToDo: Remove fruit from the board after being eaten
					delete(brd.Fruits, types.Point{X: x, Y: y})
					brd.Score++
				}
				fmt.Print("O")
			} else if _, exists := brd.Fruits[types.Point{X: x, Y: y}]; exists {
				fmt.Print("F")
			} else {
				fmt.Print(" ")
			}
		}
		println()
	}
}
