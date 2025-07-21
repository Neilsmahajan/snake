package snake

import "github.com/neilsmahajan/snake/internal/board"

func MoveSnake(snakePositionX, snakePositionY int, boardDimensions board.BoardDimensions, direction string) (int, int, bool) {
	switch direction {
	case "up":
		snakePositionY--
	case "down":
		snakePositionY++
	case "left":
		snakePositionX--
	case "right":
		snakePositionX++
	case "still":
		// Do nothing, snake stays in the same position
	}
	if snakePositionX <= 0 || snakePositionX >= boardDimensions.Width-1 || snakePositionY <= 0 || snakePositionY >= boardDimensions.Height-1 {
		return snakePositionX, snakePositionY, false // Game over if the snake hits the wall
	}
	return snakePositionX, snakePositionY, true
}
