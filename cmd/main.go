package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/neilsmahajan/snake/internal/board"
	"github.com/neilsmahajan/snake/internal/fruit"
	"github.com/neilsmahajan/snake/internal/input"
	"github.com/neilsmahajan/snake/internal/snake"
	"github.com/neilsmahajan/snake/internal/types"
)

const terminalResetDelay = 50 * time.Millisecond

var speed int

var gamePlaying = true

// resetTerminal ensures the terminal is reset to normal mode
func resetTerminal() {
	// Only reset if we're actually on a terminal
	if _, err := os.Stat("/dev/tty"); err == nil {
		// Reset terminal to normal mode
		cmd := exec.Command("stty", "sane")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run() // Ignore error as this is best effort cleanup
	}
}

func main() {
	// Ensure terminal is reset when program exits
	defer resetTerminal()

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
	stopChannel := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		input.ListenForInput(inputChannel, s, stopChannel)
	}()
	ticker := time.NewTicker(time.Duration(speed) * time.Millisecond)
	defer ticker.Stop()
	defer func() {
		close(stopChannel)             // Signal the goroutine to stop
		wg.Wait()                      // Wait for the goroutine to finish
		time.Sleep(terminalResetDelay) // Give the terminal time to reset
	}()

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
		case <-ticker.C:
			if !gamePlaying {
				break
			}
			gamePlaying = snake.MoveSnake(&brd, s)
		}
	}
	fmt.Println("Game Over! Thanks for playing!")
	if brd.Score > 0 {
		fmt.Printf("Your score: %d\n", brd.Score)
	} else {
		fmt.Println("You didn't score any points.")
	}
}
