package main

import (
	"fmt"
	"time"

	"github.com/neilsmahajan/snake/internal/board"
	"github.com/neilsmahajan/snake/internal/input"
	"github.com/neilsmahajan/snake/internal/snake"
)

var (
	width  int
	height int
)

var speed int

var (
	snakePositionX int
	snakePositionY int
)

var direction = "still"

var gamePlaying = true

func main() {
	var err error
	width, height, speed, err = input.GetDifficultyInput()
	if err != nil {
		fmt.Printf("Error getting difficulty input: %v\n", err)
		return
	}
	snakePositionX = width / 2
	snakePositionY = height / 2

	inputChannel := make(chan input.UserInput)
	go input.ListenForInput(inputChannel, &direction)
	ticker := time.NewTicker(time.Duration(speed) * time.Millisecond)
	defer ticker.Stop()

	for gamePlaying {
		board.DrawBoard(width, height, snakePositionX, snakePositionY)

		select {
		case userInput := <-inputChannel:
			if userInput.Error != nil {
				fmt.Printf("Error reading input: %v\n", userInput.Error)
				gamePlaying = false
				continue
			}
			if !userInput.GamePlaying {
				gamePlaying = false
				continue
			}
			direction = userInput.Direction
		default:
			// Do nothing, keep the snake moving
		}

		if <-ticker.C; !gamePlaying {
			break
		}

		snakePositionX, snakePositionY, gamePlaying = snake.MoveSnake(snakePositionX, snakePositionY, width, height, direction)
		// time.Sleep(200 * time.Millisecond)
	}
}
