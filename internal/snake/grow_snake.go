package snake

import "github.com/neilsmahajan/snake/internal/board"

func GrowSnake(s *Snake) {
	// Add a new segment to the snake's body at the current head position
	head := s.Body.Front().Value.(board.SnakePoint)
	newSegment := board.SnakePoint{SnakePositionX: head.SnakePositionX, SnakePositionY: head.SnakePositionY}
	s.Body.PushFront(newSegment)

	// Update the occupied map with the new segment
	s.OccupiedMap[newSegment] = s.Body.Front()

	// Note: The tail is not removed, so the snake grows longer
}
