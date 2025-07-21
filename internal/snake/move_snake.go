package snake

import (
	"github.com/neilsmahajan/snake/internal/board"
)

type Snake struct {
	SnakePoints []board.SnakePoint
	Direction   string
}

func MoveSnake(boardDimensions board.BoardDimensions, s *Snake) bool {
	switch s.Direction {
	case "up":
		s.SnakePoints[0].SnakePositionY--
	case "down":
		s.SnakePoints[0].SnakePositionY++
	case "left":
		s.SnakePoints[0].SnakePositionX--
	case "right":
		s.SnakePoints[0].SnakePositionX++
	case "still":
		// Do nothing, snake stays in the same position
	}
	if s.SnakePoints[0].SnakePositionX <= 0 || s.SnakePoints[0].SnakePositionX >= boardDimensions.Width-1 || s.SnakePoints[0].SnakePositionY <= 0 || s.SnakePoints[0].SnakePositionY >= boardDimensions.Height-1 {
		return false // Game over if the snake hits the wall
	}
	return true
}
