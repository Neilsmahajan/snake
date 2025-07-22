package board

import (
	"fmt"

	"github.com/neilsmahajan/snake/internal/fruit"
	"github.com/neilsmahajan/snake/internal/snake"
	"github.com/neilsmahajan/snake/internal/types"
)

func DrawBoard(brd *types.Board, s *types.Snake) {
	fmt.Print("\033[H\033[2J") // Clear the console
	for y := range brd.Height {
		for x := range brd.Width {
			if x == 0 || x == brd.Width-1 || y == 0 || y == brd.Height-1 {
				fmt.Print("#")
			} else if _, exists := s.OccupiedMap[types.Point{X: x, Y: y}]; exists {
				// Check to see if snake eats fruit
				if _, exists := brd.Fruits[types.Point{X: x, Y: y}]; exists {
					delete(brd.Fruits, types.Point{X: x, Y: y})
					snake.GrowSnake(s)
					fruit.CreateFruit(brd, s.OccupiedMap)
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
