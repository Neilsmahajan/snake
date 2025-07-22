package snake

import "github.com/neilsmahajan/snake/internal/types"

func GrowSnake(s *types.Snake) {
	// Add a new segment to the snake's body at the current head position
	head := s.Body.Front().Value.(types.Point)
	newSegment := types.Point{X: head.X, Y: head.Y}
	s.Body.PushFront(newSegment)

	// Update the occupied map with the new segment
	s.OccupiedMap[newSegment] = s.Body.Front()

	// Note: The tail is not removed, so the snake grows longer
}
