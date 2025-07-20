package main

import (
	"time"

	"github.com/neilsmahajan/snake/internal/board"
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
		snakePositionX, snakePositionY, gameOver = snake.MoveSnake(snakePositionX, snakePositionY, width, height, direction)
		// Simulate a delay for the snake movement
		time.Sleep(200 * time.Millisecond)
	}
}
