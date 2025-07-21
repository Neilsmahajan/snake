package snake

import (
	"github.com/neilsmahajan/snake/internal/board"
)

func MoveSnake(brd board.Board, s *Snake) bool {
	if s.Direction == "still" {
		return true // No movement, just return
	}

	head := s.Body.Front().Value.(board.SnakePoint)
	var newHead board.SnakePoint

	switch s.Direction {
	case "up":
		newHead = board.SnakePoint{SnakePositionX: head.SnakePositionX, SnakePositionY: head.SnakePositionY - 1}
	case "down":
		newHead = board.SnakePoint{SnakePositionX: head.SnakePositionX, SnakePositionY: head.SnakePositionY + 1}
	case "left":
		newHead = board.SnakePoint{SnakePositionX: head.SnakePositionX - 1, SnakePositionY: head.SnakePositionY}
	case "right":
		newHead = board.SnakePoint{SnakePositionX: head.SnakePositionX + 1, SnakePositionY: head.SnakePositionY}
	default:
		return false // Invalid direction
	}

	if newHead.SnakePositionX <= 0 || newHead.SnakePositionX >= brd.Width-1 ||
		newHead.SnakePositionY <= 0 || newHead.SnakePositionY >= brd.Height-1 {
		return false // Hit the wall
	}

	if _, exists := s.OccupiedMap[newHead]; exists {
		return false // Hit itself
	}

	s.Body.PushFront(newHead)
	s.OccupiedMap[newHead] = s.Body.Front()

	if s.Body.Len() > 1 {
		tail := s.Body.Back()
		delete(s.OccupiedMap, tail.Value.(board.SnakePoint))
		s.Body.Remove(tail)
	}

	return true
}
