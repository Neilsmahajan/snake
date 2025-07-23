package snake

import (
	"testing"

	"github.com/neilsmahajan/snake/internal/types"
)

func TestNewSnake(t *testing.T) {
	tests := []struct {
		name    string
		board   types.Board
		wantX   int
		wantY   int
		wantLen int
	}{
		{
			name:    "Small board",
			board:   types.Board{Width: 20, Height: 10},
			wantX:   10,
			wantY:   5,
			wantLen: 1,
		},
		{
			name:    "Medium board",
			board:   types.Board{Width: 40, Height: 20},
			wantX:   20,
			wantY:   10,
			wantLen: 1,
		},
		{
			name:    "Large board",
			board:   types.Board{Width: 80, Height: 40},
			wantX:   40,
			wantY:   20,
			wantLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			snake := NewSnake(tt.board)

			// Test snake initialization
			if snake == nil {
				t.Error("NewSnake() returned nil")
				return
			}

			// Test initial body length
			if snake.Body.Len() != tt.wantLen {
				t.Errorf("Snake body length = %d, want %d", snake.Body.Len(), tt.wantLen)
			}

			// Test initial position
			if snake.Body.Len() > 0 {
				head := snake.Body.Front().Value.(types.Point)
				if head.X != tt.wantX {
					t.Errorf("Snake head X = %d, want %d", head.X, tt.wantX)
				}
				if head.Y != tt.wantY {
					t.Errorf("Snake head Y = %d, want %d", head.Y, tt.wantY)
				}
			}

			// Test initial direction
			if snake.Direction != "still" {
				t.Errorf("Snake initial direction = %s, want 'still'", snake.Direction)
			}

			// Test initial growth state
			if snake.ShouldGrow {
				t.Error("Snake should not grow initially")
			}

			// Test occupied map
			if len(snake.OccupiedMap) != tt.wantLen {
				t.Errorf("OccupiedMap length = %d, want %d", len(snake.OccupiedMap), tt.wantLen)
			}

			// Test that head is in occupied map
			if snake.Body.Len() > 0 {
				head := snake.Body.Front().Value.(types.Point)
				if _, exists := snake.OccupiedMap[head]; !exists {
					t.Error("Snake head not found in OccupiedMap")
				}
			}
		})
	}
}

func TestNewSnakeConsistency(t *testing.T) {
	board := types.Board{Width: 20, Height: 10}

	// Create multiple snakes and ensure they're all at the same position
	snakes := make([]*types.Snake, 10)
	for i := range snakes {
		snakes[i] = NewSnake(board)
	}

	firstHead := snakes[0].Body.Front().Value.(types.Point)
	for i := 1; i < len(snakes); i++ {
		head := snakes[i].Body.Front().Value.(types.Point)
		if head.X != firstHead.X || head.Y != firstHead.Y {
			t.Errorf("Snake %d head position (%d,%d) differs from expected (%d,%d)",
				i, head.X, head.Y, firstHead.X, firstHead.Y)
		}
	}
}

func TestSnakeDataStructureIntegrity(t *testing.T) {
	board := types.Board{Width: 20, Height: 10}
	snake := NewSnake(board)

	// Test that Body and OccupiedMap are in sync
	bodyCount := snake.Body.Len()
	mapCount := len(snake.OccupiedMap)

	if bodyCount != mapCount {
		t.Errorf("Body length (%d) != OccupiedMap length (%d)", bodyCount, mapCount)
	}

	// Iterate through body and verify each segment is in the map
	for e := snake.Body.Front(); e != nil; e = e.Next() {
		point := e.Value.(types.Point)
		if element, exists := snake.OccupiedMap[point]; !exists {
			t.Errorf("Point (%d,%d) in body but not in OccupiedMap", point.X, point.Y)
		} else if element != e {
			t.Errorf("OccupiedMap points to wrong element for point (%d,%d)", point.X, point.Y)
		}
	}

	// Verify each map entry points to a valid body segment
	for point, element := range snake.OccupiedMap {
		if element.Value.(types.Point) != point {
			t.Errorf("OccupiedMap entry for (%d,%d) points to wrong body segment", point.X, point.Y)
		}
	}
}

// Benchmark for snake creation
func BenchmarkNewSnake(b *testing.B) {
	board := types.Board{Width: 40, Height: 20}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewSnake(board)
	}
}

// Benchmark for different board sizes
func BenchmarkNewSnakeSmall(b *testing.B) {
	board := types.Board{Width: 20, Height: 10}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewSnake(board)
	}
}

func BenchmarkNewSnakeLarge(b *testing.B) {
	board := types.Board{Width: 80, Height: 40}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewSnake(board)
	}
}

// Test memory usage
func TestSnakeMemoryUsage(t *testing.T) {
	board := types.Board{Width: 20, Height: 10}
	snake := NewSnake(board)

	// Verify minimal memory usage for initial snake
	if len(snake.OccupiedMap) > 1 {
		t.Errorf("Initial snake using too much memory: %d map entries", len(snake.OccupiedMap))
	}

	if snake.Body.Len() > 1 {
		t.Errorf("Initial snake body too long: %d segments", snake.Body.Len())
	}
}
