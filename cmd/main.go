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

var gamePlaying = true

func main() {
	for gamePlaying {
		board.DrawBoard(width, height, snakePositionX, snakePositionY)
		var err error
		direction, gamePlaying, err = input.GetUserInput(direction)
		if err != nil {
			panic(err)
		}
		if !gamePlaying {
			break
		}
		snakePositionX, snakePositionY, gamePlaying = snake.MoveSnake(snakePositionX, snakePositionY, width, height, direction)
		time.Sleep(200 * time.Millisecond)
	}
}
