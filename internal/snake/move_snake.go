package snake

import (
	"github.com/neilsmahajan/snake/internal/fruit"
	"github.com/neilsmahajan/snake/internal/types"
)

func MoveSnake(brd *types.Board, s *types.Snake) bool {
	if s.Direction == types.DirectionStill {
		return true // No movement, just return
	}

	head := s.Body.Front().Value.(types.Point)
	var newHead types.Point

	switch s.Direction {
	case types.DirectionUp:
		newHead = types.Point{X: head.X, Y: head.Y - 1}
	case types.DirectionDown:
		newHead = types.Point{X: head.X, Y: head.Y + 1}
	case types.DirectionLeft:
		newHead = types.Point{X: head.X - 1, Y: head.Y}
	case types.DirectionRight:
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

	// Check if snake eats fruit
	if _, exists := brd.Fruits[newHead]; exists {
		delete(brd.Fruits, newHead)
		fruit.CreateFruit(brd, s.OccupiedMap)
		brd.Score++
		s.ShouldGrow = true
	}

	s.Body.PushFront(newHead)
	s.OccupiedMap[newHead] = s.Body.Front()

	// Only remove tail if snake shouldn't grow
	if s.ShouldGrow {
		s.ShouldGrow = false // Reset the grow flag
	} else if s.Body.Len() > 1 {
		tail := s.Body.Back()
		delete(s.OccupiedMap, tail.Value.(types.Point))
		s.Body.Remove(tail)
	}

	return true
}
