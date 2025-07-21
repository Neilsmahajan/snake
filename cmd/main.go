package main

import (
	"fmt"
	"time"

	"github.com/neilsmahajan/snake/internal/board"
	"github.com/neilsmahajan/snake/internal/input"
	"github.com/neilsmahajan/snake/internal/snake"
)

var speed int

var gamePlaying = true

func main() {
	var err error
	var boardDimensions board.BoardDimensions
	boardDimensions, speed, err = input.GetDifficultyInput()
	if err != nil {
		fmt.Printf("Error getting difficulty input: %v\n", err)
		return
	}

	s := snake.NewSnake(boardDimensions)

	inputChannel := make(chan input.UserInput)
	go input.ListenForInput(inputChannel, s)
	ticker := time.NewTicker(time.Duration(speed) * time.Millisecond)
	defer ticker.Stop()

	for gamePlaying {
		board.DrawBoard(boardDimensions, s.OccupiedMap)

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
			s.Direction = userInput.Direction
		default:
			// Do nothing, keep the snake moving
		}

		if <-ticker.C; !gamePlaying {
			break
		}

		gamePlaying = snake.MoveSnake(boardDimensions, s)
	}
}
