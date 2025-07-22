package fruit

import (
	"container/list"
	"math/rand/v2"

	"github.com/neilsmahajan/snake/internal/types"
)

const borderOffset = 2

func CreateFruit(brd *types.Board, occupiedMap map[types.Point]*list.Element) {
	// Create a new fruit at a random position on the board that is not occupied by the snake
	var newFruit types.Point
	for {
		newFruit = types.Point{
			X: rand.IntN(brd.Width-borderOffset) + 1,  // Avoid borders
			Y: rand.IntN(brd.Height-borderOffset) + 1, // Avoid borders
		}
		// Check if the new fruit is not occupied by the snake and not already in a fruit position
		_, existsInSnake := occupiedMap[newFruit]
		_, existsInFruits := brd.Fruits[newFruit]
		if !existsInSnake && !existsInFruits {
			break
		}
	}

	// Add the new fruit to the board's fruit list
	brd.Fruits[newFruit] = struct{}{}
}
