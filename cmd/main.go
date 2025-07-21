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
	fmt.Println("Welcome to the Snake Game!")
	fmt.Println("Use 'w' or 'k' to move up, 's' or 'j' to move down, 'a' or 'h' to move left, 'd' or 'l' to move right.")
	fmt.Println("Press 'q' or 'Esc' to quit the game.")

	inputChannel := make(chan input.UserInput)
	go input.ListenForInput(inputChannel)

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

		snakePositionX, snakePositionY, gamePlaying = snake.MoveSnake(snakePositionX, snakePositionY, width, height, direction)
		time.Sleep(200 * time.Millisecond)
	}
}
