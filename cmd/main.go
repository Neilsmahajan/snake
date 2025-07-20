package main

import (
	"time"

	"github.com/neilsmahajan/snake/internal/board"
	"github.com/neilsmahajan/snake/internal/input"
	"github.com/neilsmahajan/snake/internal/snake"
)

const (
	width  = 20
	height = 10
)

var (
	snakePositionX = width / 2
	snakePositionY = height / 2
)

var direction = "right"

var gameOver = false

func main() {
	for !gameOver {
		board.DrawBoard(width, height, snakePositionX, snakePositionY)
		direction, gameOver = input.GetUserInput(direction)
		snakePositionX, snakePositionY, gameOver = snake.MoveSnake(snakePositionX, snakePositionY, width, height, direction)
		time.Sleep(200 * time.Millisecond)
	}
}
