package snake

import (
	"container/list"

	"github.com/neilsmahajan/snake/internal/board"
)

type Snake struct {
	Body        *list.List                         // head = Front(), tail = Back()
	OccupiedMap map[board.SnakePoint]*list.Element // Maps snake points to their list elements for quick access
	Direction   string
}

func NewSnake(brd board.Board) *Snake {
	s := &Snake{
		Body:        list.New(),
		OccupiedMap: make(map[board.SnakePoint]*list.Element),
		Direction:   "still",
	}
	start := board.SnakePoint{SnakePositionX: brd.Width / 2, SnakePositionY: brd.Height / 2}
	e := s.Body.PushFront(start)
	s.OccupiedMap[start] = e
	return s
}
