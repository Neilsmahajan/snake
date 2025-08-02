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

const (
	terminalResetDelay = 100 * time.Millisecond
)

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

// setupGame initializes the game board and snake
func setupGame() (types.Board, *types.Snake, error) {
	brd, gameSpeed, err := input.GetDifficultyInput()
	if err != nil {
		return types.Board{}, nil, fmt.Errorf("error getting difficulty input: %v", err)
	}
	speed = gameSpeed

	s := snake.NewSnake(brd)
	fruit.CreateFruit(&brd, s.OccupiedMap)

	return brd, s, nil
}

// runGameLoop runs the main game loop
func runGameLoop(brd *types.Board, s *types.Snake) {
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
		board.DrawBoard(brd, s)

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
			gamePlaying = snake.MoveSnake(brd, s)
		}
	}
}

// showGameOverScreen displays the final score and game over message
func showGameOverScreen(brd *types.Board) {
	// Clear screen and show final message
	fmt.Print("\033[H\033[2J")

	// Game Over banner
	fmt.Printf("%s%s", types.ColorBold, types.ColorRed)
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║             GAME OVER!               ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Printf("%s\n", types.ColorReset)

	// Score display
	if brd.Score > 0 {
		fmt.Printf("%s🏆 Final Score: %s%s%d%s%s\n\n",
			types.ColorYellow, types.ColorBold, types.ColorGreen, brd.Score, types.ColorReset, types.ColorYellow)

		// Score rating
		switch {
		case brd.Score >= 50:
			fmt.Printf("%s🌟 Amazing! You're a Snake Master! 🌟%s\n", types.ColorGreen, types.ColorReset)
		case brd.Score >= 30:
			fmt.Printf("%s⭐ Great job! You're getting good! ⭐%s\n", types.ColorCyan, types.ColorReset)
		case brd.Score >= 15:
			fmt.Printf("%s✨ Not bad! Keep practicing! ✨%s\n", types.ColorYellow, types.ColorReset)
		case brd.Score >= 5:
			fmt.Printf("%s🎯 Good start! Try again! 🎯%s\n", types.ColorBlue, types.ColorReset)
		default:
			fmt.Printf("%s🎮 Keep trying! You'll get better! 🎮%s\n", types.ColorPurple, types.ColorReset)
		}
	} else {
		fmt.Printf("%s😔 No points scored this time!%s\n", types.ColorRed, types.ColorReset)
		fmt.Printf("%s💪 Don't give up - try again! 💪%s\n", types.ColorGreen, types.ColorReset)
	}

	fmt.Printf("\n%sThanks for playing Snake! 🐍%s\n", types.ColorCyan, types.ColorReset)
	fmt.Printf("%sPress %sENTER%s to exit...%s\n", types.ColorWhite, types.ColorGreen, types.ColorWhite, types.ColorReset)

	// Wait for Enter key to exit gracefully
	var input string
	_, _ = fmt.Scanln(&input) // Explicitly ignore both return values to satisfy linters
}

func main() {
	// Ensure terminal is reset when program exits (multiple defers for safety)
	defer func() {
		resetTerminal()
		time.Sleep(terminalResetDelay)
	}()
	defer resetTerminal()

	// Setup the game
	brd, s, err := setupGame()
	if err != nil {
		fmt.Printf("Error setting up game: %v\n", err)
		return
	}

	// Run the main game loop
	runGameLoop(&brd, s)

	// Show game over screen
	showGameOverScreen(&brd)
}
