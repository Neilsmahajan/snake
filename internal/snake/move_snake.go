package snake

import (
	"github.com/neilsmahajan/snake/internal/board"
)

type Snake struct {
	Body      []board.SnakePoint
	Direction string
}

func updateSnakeBody(s *Snake) bool {
	if len(s.Body) == 1 {
		return true
	}
	// Check for self-collision
	head := s.Body[len(s.Body)-1]
	for i := 0; i < len(s.Body)-1; i++ {
		if s.Body[i].SnakePositionX == head.SnakePositionX && s.Body[i].SnakePositionY == head.SnakePositionY {
			return false // Game over if the snake collides with itself
		}
	}
	for i := len(s.Body) - 2; i > 0; i-- {
		s.Body[i].SnakePositionX = s.Body[i-1].SnakePositionX
		s.Body[i].SnakePositionY = s.Body[i-1].SnakePositionY
	}

	return true
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
	if !updateSnakeBody(s) {
		return false // Game over if the snake collides with itself
	}
	return true
}
