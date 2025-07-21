package snake

import (
	"github.com/neilsmahajan/snake/internal/board"
)

type Snake struct {
	Body      []board.SnakePoint
	Direction string
}

func MoveSnake(boardDimensions board.BoardDimensions, s *Snake) bool {
	head := &s.Body[len(s.Body)-1]
	switch s.Direction {
	case "up":
		head.SnakePositionY--
	case "down":
		head.SnakePositionY++
	case "left":
		head.SnakePositionX--
	case "right":
		head.SnakePositionX++
	case "still":
		return true // No movement, so the snake is still valid
	}
	if head.SnakePositionX <= 0 || head.SnakePositionX >= boardDimensions.Width-1 || head.SnakePositionY <= 0 || head.SnakePositionY >= boardDimensions.Height-1 {
		return false // Game over if the snake hits the wall
	}
	return true
}
