package fruit

import (
	"container/list"
	"crypto/rand"
	"math/big"

	"github.com/neilsmahajan/snake/internal/types"
)

const borderOffset = 2

func CreateFruit(brd *types.Board, occupiedMap map[types.Point]*list.Element) {
	// Create a new fruit at a random position on the board that is not occupied by the snake
	var newFruit types.Point
	for {
		// Generate cryptographically secure random numbers
		xRange := big.NewInt(int64(brd.Width - borderOffset))
		yRange := big.NewInt(int64(brd.Height - borderOffset))

		xRand, err := rand.Int(rand.Reader, xRange)
		if err != nil {
			// Fallback to a simple position if crypto/rand fails
			newFruit = types.Point{X: 1, Y: 1}
		} else {
			yRand, err := rand.Int(rand.Reader, yRange)
			if err != nil {
				// Fallback to a simple position if crypto/rand fails
				newFruit = types.Point{X: 1, Y: 1}
			} else {
				newFruit = types.Point{
					X: int(xRand.Int64()) + 1, // Avoid borders
					Y: int(yRand.Int64()) + 1, // Avoid borders
				}
			}
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
