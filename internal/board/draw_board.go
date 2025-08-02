package board

import (
	"fmt"

	"github.com/neilsmahajan/snake/internal/types"
)

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
	fmt.Print(types.ClearScreen) // Clear the console

	// Display score and game info
	fmt.Printf("%s%s🐍 SNAKE GAME 🐍%s\n", types.ColorBold, types.ColorGreen, types.ColorReset)
	fmt.Printf("%sScore: %s%d%s\n", types.ColorCyan, types.ColorYellow, brd.Score, types.ColorReset)
	fmt.Println()

	for y := range brd.Height {
		for x := range brd.Width {
			cellType := getCellType(x, y, brd, s)
			switch cellType {
			case types.CellWall:
				fmt.Printf("%s%s%s", types.ColorBlue, types.WallSymbol, types.ColorReset)
			case types.CellSnake:
				fmt.Printf("%s%s%s", types.ColorGreen, types.SnakeSymbol, types.ColorReset)
			case types.CellFruit:
				fmt.Printf("%s%s%s", types.ColorRed, types.FruitSymbol, types.ColorReset)
			case types.CellEmpty:
				fmt.Print(types.EmptySymbol)
			}
		}
		println()
	}

	// Display controls at the bottom
	fmt.Printf("\n%s%sControls:%s %sW/K↑ S/J↓ A/H← D/L→ Q/ESC=Quit%s\n",
		types.ColorBold, types.ColorCyan, types.ColorReset, types.ColorWhite, types.ColorReset)
}
