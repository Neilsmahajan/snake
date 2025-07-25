package types

import "container/list"

// Direction constants
const (
	DirectionUp    = "up"
	DirectionDown  = "down"
	DirectionLeft  = "left"
	DirectionRight = "right"
	DirectionStill = "still"
)

// Point represents a coordinate on the game board
type Point struct {
	X int
	Y int
}

// Board represents the game board
type Board struct {
	Width  int
	Height int
	Fruits map[Point]struct{}
	Score  int
}

// Snake represents the snake in the game
type Snake struct {
	Body        *list.List              // head = Front(), tail = Back()
	OccupiedMap map[Point]*list.Element // Maps snake points to their list elements for quick access
	Direction   string
	ShouldGrow  bool // Flag to indicate if snake should grow on next move
}

// UserInput represents input from the user
type UserInput struct {
	Direction   string
	GamePlaying bool
	Error       error
}
