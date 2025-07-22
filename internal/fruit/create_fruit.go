package fruit

import (
	"container/list"
	"math/rand"

	"github.com/neilsmahajan/snake/internal/board"
)

func CreateFruit(brd *board.Board, occupiedMap map[board.SnakePoint]*list.Element) {
	// Create a new fruit at a random position on the board that is not occupied by the snake
	var newFruit board.FruitCoordinate
	for {
		newFruit = board.FruitCoordinate{
			FruitPositionX: rand.Intn(brd.Width),
			FruitPositionY: rand.Intn(brd.Height),
		}
		// Check if the new fruit is not occupied by the snake
		if _, exists := occupiedMap[board.SnakePoint{SnakePositionX: newFruit.FruitPositionX, SnakePositionY: newFruit.FruitPositionY}]; !exists {
			break
		}
	}

	// Add the new fruit to the board's fruit list
	brd.Fruits[newFruit] = struct{}{}
}
