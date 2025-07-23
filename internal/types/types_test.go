package types

import (
	"container/list"
	"testing"
)

func TestPoint(t *testing.T) {
	tests := []struct {
		name string
		p    Point
		x    int
		y    int
	}{
		{
			name: "Zero point",
			p:    Point{X: 0, Y: 0},
			x:    0,
			y:    0,
		},
		{
			name: "Positive coordinates",
			p:    Point{X: 5, Y: 10},
			x:    5,
			y:    10,
		},
		{
			name: "Negative coordinates",
			p:    Point{X: -3, Y: -7},
			x:    -3,
			y:    -7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.p.X != tt.x {
				t.Errorf("Point.X = %d, want %d", tt.p.X, tt.x)
			}
			if tt.p.Y != tt.y {
				t.Errorf("Point.Y = %d, want %d", tt.p.Y, tt.y)
			}
		})
	}
}

func TestBoard(t *testing.T) {
	board := Board{
		Width:  20,
		Height: 10,
		Fruits: make(map[Point]struct{}),
		Score:  0,
	}

	if board.Width != 20 {
		t.Errorf("Board.Width = %d, want 20", board.Width)
	}
	if board.Height != 10 {
		t.Errorf("Board.Height = %d, want 10", board.Height)
	}
	if board.Score != 0 {
		t.Errorf("Board.Score = %d, want 0", board.Score)
	}
	if len(board.Fruits) != 0 {
		t.Errorf("Board.Fruits length = %d, want 0", len(board.Fruits))
	}
}

func TestSnake(t *testing.T) {
	snake := Snake{
		Body:        list.New(),
		OccupiedMap: make(map[Point]*list.Element),
		Direction:   DirectionUp,
		ShouldGrow:  false,
	}

	if snake.Body.Len() != 0 {
		t.Errorf("Snake.Body.Len() = %d, want 0", snake.Body.Len())
	}
	if len(snake.OccupiedMap) != 0 {
		t.Errorf("Snake.OccupiedMap length = %d, want 0", len(snake.OccupiedMap))
	}
	if snake.Direction != DirectionUp {
		t.Errorf("Snake.Direction = %s, want 'up'", snake.Direction)
	}
	if snake.ShouldGrow {
		t.Errorf("Snake.ShouldGrow = %t, want false", snake.ShouldGrow)
	}
}

func TestUserInput(t *testing.T) {
	tests := []struct {
		name      string
		input     UserInput
		direction string
		playing   bool
	}{
		{
			name: "Valid input",
			input: UserInput{
				Direction:   DirectionUp,
				GamePlaying: true,
				Error:       nil,
			},
			direction: DirectionUp,
			playing:   true,
		},
		{
			name: "Game stopped",
			input: UserInput{
				Direction:   "",
				GamePlaying: false,
				Error:       nil,
			},
			direction: "",
			playing:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input.Direction != tt.direction {
				t.Errorf("UserInput.Direction = %s, want %s", tt.input.Direction, tt.direction)
			}
			if tt.input.GamePlaying != tt.playing {
				t.Errorf("UserInput.GamePlaying = %t, want %t", tt.input.GamePlaying, tt.playing)
			}
		})
	}
}

// Benchmark for Point creation
func BenchmarkPointCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Point{X: i, Y: i * 2}
	}
}

// Benchmark for map operations with Point as key
func BenchmarkPointMapOperations(b *testing.B) {
	m := make(map[Point]struct{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := Point{X: i % 100, Y: (i * 2) % 100}
		m[p] = struct{}{}
		_, exists := m[p]
		if !exists {
			b.Error("Point should exist in map")
		}
	}
}
