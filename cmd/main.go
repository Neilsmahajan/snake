package main

import (
	"github.com/neilsmahajan/snake/internal/board"
)

const (
	width  = 20
	height = 10
)

var (
	snakePositionX = width / 2
	snakePositionY = height / 2
)

func main() {
	board.DrawBoard(width, height, snakePositionX, snakePositionY)
}
