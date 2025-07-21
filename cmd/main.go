package main

import (
	"fmt"
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

var direction = "still"

var gamePlaying = true

func main() {
	for gamePlaying {
		board.DrawBoard(width, height, snakePositionX, snakePositionY)
		inputChannel := make(chan input.UserInput)
		go input.GetUserInput(direction, inputChannel)
		userInput := <-inputChannel
		if userInput.Error != nil {
			fmt.Printf("Error reading input: %v\n", userInput.Error)
		}
		if !userInput.GamePlaying {
			gamePlaying = false
			continue
		}
		direction = userInput.Direction
		//if !gamePlaying {
		//	break
		//}
		snakePositionX, snakePositionY, gamePlaying = snake.MoveSnake(snakePositionX, snakePositionY, width, height, direction)
		time.Sleep(200 * time.Millisecond)
	}
}
