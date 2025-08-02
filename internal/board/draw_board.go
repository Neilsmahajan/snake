package board

import (
	"fmt"

	"github.com/neilsmahajan/snake/internal/types"
)

const clearScreenSequence = "\033[H\033[2J"

// getCellType determines what type of content is at the given position
func getCellType(x, y int, brd *types.Board, s *types.Snake) types.CellType {
	// Check for wall (border)
	if x == 0 || x == brd.Width-1 || y == 0 || y == brd.Height-1 {
		return types.CellWall
	}

	point := types.Point{X: x, Y: y}

	// Check for snake
	if _, exists := s.OccupiedMap[point]; exists {
		return types.CellSnake
	}

	// Check for fruit
	if _, exists := brd.Fruits[point]; exists {
		return types.CellFruit
	}

	// Default to empty
	return types.CellEmpty
}

func DrawBoard(brd *types.Board, s *types.Snake) {
	fmt.Print(clearScreenSequence) // Clear the console
	for y := range brd.Height {
		for x := range brd.Width {
			cellType := getCellType(x, y, brd, s)
			switch cellType {
			case types.CellWall:
				fmt.Print("#")
			case types.CellSnake:
				fmt.Print("O")
			case types.CellFruit:
				fmt.Print("F")
			case types.CellEmpty:
				fmt.Print(" ")
			}
		}
		println()
	}
}
