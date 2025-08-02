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

const terminalResetDelay = 100 * time.Millisecond

var speed int

var gamePlaying = true

// resetTerminal ensures the terminal is reset to normal mode
func resetTerminal() {
	// Multiple cleanup attempts to ensure terminal is properly reset
	// Using explicit commands to satisfy gosec G204 requirements

	// Reset terminal to sane state
	cmd1 := exec.Command("stty", "sane") // #nosec G204 - command is hardcoded
	cmd1.Stdin = os.Stdin
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr
	_ = cmd1.Run() // Ignore errors as this is best effort cleanup

	// Enable echo
	cmd2 := exec.Command("stty", "echo") // #nosec G204 - command is hardcoded
	cmd2.Stdin = os.Stdin
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	_ = cmd2.Run() // Ignore errors as this is best effort cleanup

	// Enable canonical mode
	cmd3 := exec.Command("stty", "icanon") // #nosec G204 - command is hardcoded
	cmd3.Stdin = os.Stdin
	cmd3.Stdout = os.Stdout
	cmd3.Stderr = os.Stderr
	_ = cmd3.Run() // Ignore errors as this is best effort cleanup

	// Force flush stdout/stderr
	_ = os.Stdout.Sync() // #nosec G104 - stdout sync, error doesn't affect cleanup
	_ = os.Stderr.Sync() // #nosec G104 - stderr sync, error doesn't affect cleanup
}

func main() {
	// Ensure terminal is reset when program exits (multiple defers for safety)
	defer func() {
		resetTerminal()
		time.Sleep(terminalResetDelay)
	}()
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

	inputChannel := make(chan types.UserInput, 10) // Buffered channel to prevent blocking
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

		// Drain any remaining input
		for len(inputChannel) > 0 {
			<-inputChannel
		}
		close(inputChannel)
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

	// Clear screen and show final message
	fmt.Print("\033[H\033[2J")
	fmt.Println("Game Over! Thanks for playing!")
	if brd.Score > 0 {
		fmt.Printf("Your score: %d\n", brd.Score)
	} else {
		fmt.Println("You didn't score any points.")
	}
	fmt.Println("Press Enter to exit...")

	// Wait for Enter key to exit gracefully
	var input string
	_, _ = fmt.Scanln(&input) // Explicitly ignore both return values to satisfy linters
}
