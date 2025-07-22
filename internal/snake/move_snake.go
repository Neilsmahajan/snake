package snake

import (
	"github.com/neilsmahajan/snake/internal/types"
)

func MoveSnake(brd types.Board, s *types.Snake) bool {
	if s.Direction == "still" {
		return true // No movement, just return
	}

	head := s.Body.Front().Value.(types.Point)
	var newHead types.Point

	switch s.Direction {
	case "up":
		newHead = types.Point{X: head.X, Y: head.Y - 1}
	case "down":
		newHead = types.Point{X: head.X, Y: head.Y + 1}
	case "left":
		newHead = types.Point{X: head.X - 1, Y: head.Y}
	case "right":
		newHead = types.Point{X: head.X + 1, Y: head.Y}
	default:
		return false // Invalid direction
	}

	if newHead.X <= 0 || newHead.X >= brd.Width-1 ||
		newHead.Y <= 0 || newHead.Y >= brd.Height-1 {
		return false // Hit the wall
	}

	if _, exists := s.OccupiedMap[newHead]; exists {
		return false // Hit itself
	}

	s.Body.PushFront(newHead)
	s.OccupiedMap[newHead] = s.Body.Front()

	if s.Body.Len() > 1 {
		tail := s.Body.Back()
		delete(s.OccupiedMap, tail.Value.(types.Point))
		s.Body.Remove(tail)
	}

	return true
}
