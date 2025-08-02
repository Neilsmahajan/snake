package types

import "container/list"

// ANSI color codes
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

// Game display constants
const (
	ClearScreen = "\033[H\033[2J"
	WallSymbol  = "█" // Solid block
	SnakeSymbol = "●" // Circle
	FruitSymbol = "♦" // Diamond (single-width character)
	EmptySymbol = " " // Space
)

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

// CellType represents the type of content in a board cell
type CellType int

const (
	CellWall CellType = iota
	CellSnake
	CellFruit
	CellEmpty
)

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
