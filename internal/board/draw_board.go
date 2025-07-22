package board

import (
	"container/list"
	"fmt"
)

type Board struct {
	Width  int
	Height int
	Fruits map[FruitCoordinate]struct{}
	Score  int
}

type SnakePoint struct {
	SnakePositionX int
	SnakePositionY int
}
type FruitCoordinate struct {
	FruitPositionX int
	FruitPositionY int
}

func DrawBoard(brd *Board, occupiedMap map[SnakePoint]*list.Element) {
	fmt.Print("\033[H\033[2J") // Clear the console
	for y := range brd.Height {
		for x := range brd.Width {
			if x == 0 || x == brd.Width-1 || y == 0 || y == brd.Height-1 {
				fmt.Print("#")
			} else if _, exists := occupiedMap[SnakePoint{SnakePositionX: x, SnakePositionY: y}]; exists {
				// Check to see if snake eats fruit
				if _, exists := brd.Fruits[FruitCoordinate{FruitPositionX: x, FruitPositionY: y}]; exists {
					// ToDo: Remove fruit from the board after being eaten
					delete(brd.Fruits, FruitCoordinate{FruitPositionX: x, FruitPositionY: y})
					brd.Score++
				}
				fmt.Print("O")
			} else if _, exists := brd.Fruits[FruitCoordinate{FruitPositionX: x, FruitPositionY: y}]; exists {
				fmt.Print("F")
			} else {
				fmt.Print(" ")
			}
		}
		println()
	}
}
