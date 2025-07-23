package snake

import (
	"container/list"

	"github.com/neilsmahajan/snake/internal/types"
)

func NewSnake(brd types.Board) *types.Snake {
	s := &types.Snake{
		Body:        list.New(),
		OccupiedMap: make(map[types.Point]*list.Element),
		Direction:   types.DirectionStill,
	}
	start := types.Point{X: brd.Width / 2, Y: brd.Height / 2}
	e := s.Body.PushFront(start)
	s.OccupiedMap[start] = e
	return s
}
