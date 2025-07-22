package main

import (
	"fmt"
	"time"

	"github.com/neilsmahajan/snake/internal/board"
	"github.com/neilsmahajan/snake/internal/fruit"
	"github.com/neilsmahajan/snake/internal/input"
	"github.com/neilsmahajan/snake/internal/snake"
	"github.com/neilsmahajan/snake/internal/types"
)

var speed int

var gamePlaying = true

func main() {
	var err error
	var brd types.Board
	brd, speed, err = input.GetDifficultyInput()
	if err != nil {
		fmt.Printf("Error getting difficulty input: %v\n", err)
		return
	}

	s := snake.NewSnake(brd)
	fruit.CreateFruit(&brd, s.OccupiedMap)

	inputChannel := make(chan types.UserInput)
	go input.ListenForInput(inputChannel, s)
	ticker := time.NewTicker(time.Duration(speed) * time.Millisecond)
	defer ticker.Stop()

	for gamePlaying {
		board.DrawBoard(&brd, s)

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

		gamePlaying = snake.MoveSnake(brd, s)
	}
	fmt.Println("Game Over! Thanks for playing!")
	if brd.Score > 0 {
		fmt.Printf("Your score: %d\n", brd.Score)
	} else {
		fmt.Println("You didn't score any points.")
	}
}
