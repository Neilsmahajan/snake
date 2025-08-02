package input

import (
	"fmt"

	"github.com/neilsmahajan/snake/internal/types"
)

// displayWelcomeScreen shows the welcome banner
func displayWelcomeScreen() {
	fmt.Print(types.ClearScreen)
	fmt.Printf("%s%s", types.ColorBold, types.ColorGreen)
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║          🐍 SNAKE GAME 🐍           ║")
	fmt.Println("║      Written in Go Language         ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Printf("%s\n", types.ColorReset)
}

// getBoardSize prompts user for board size and returns width, height
func getBoardSize() (width, height int, err error) {
	fmt.Printf("%s%sBoard Size Selection:%s\n", types.ColorBold, types.ColorCyan, types.ColorReset)
	fmt.Printf("  %s[s]%s - Small  (20x10) 📱\n", types.ColorYellow, types.ColorReset)
	fmt.Printf("  %s[m]%s - Medium (40x20) 💻\n", types.ColorYellow, types.ColorReset)
	fmt.Printf("  %s[l]%s - Large  (80x40) 🖥️\n", types.ColorYellow, types.ColorReset)
	fmt.Printf("%sEnter your choice: %s", types.ColorWhite, types.ColorReset)

	var size string
	if _, err = fmt.Scanln(&size); err != nil {
		return 0, 0, fmt.Errorf("error reading size input: %v", err)
	}

	switch size {
	case "s":
		fmt.Printf("%s✓ Small board selected!%s\n\n", types.ColorGreen, types.ColorReset)
		return 20, 10, nil
	case "m":
		fmt.Printf("%s✓ Medium board selected!%s\n\n", types.ColorGreen, types.ColorReset)
		return 40, 20, nil
	case "l":
		fmt.Printf("%s✓ Large board selected!%s\n\n", types.ColorGreen, types.ColorReset)
		return 80, 40, nil
	default:
		return 0, 0, fmt.Errorf("invalid size input: %s", size)
	}
}

// getGameSpeed prompts user for game speed and returns speed in milliseconds
func getGameSpeed() (int, error) {
	fmt.Printf("%s%sSpeed Selection:%s\n", types.ColorBold, types.ColorCyan, types.ColorReset)
	fmt.Printf("  %s[s]%s - Slow   (200ms) 🐌\n", types.ColorYellow, types.ColorReset)
	fmt.Printf("  %s[m]%s - Medium (100ms) 🚶\n", types.ColorYellow, types.ColorReset)
	fmt.Printf("  %s[f]%s - Fast   (50ms)  🏃\n", types.ColorYellow, types.ColorReset)
	fmt.Printf("%sEnter your choice: %s", types.ColorWhite, types.ColorReset)

	var speedInput string
	if _, err := fmt.Scanln(&speedInput); err != nil {
		return 0, fmt.Errorf("error reading speed input: %v", err)
	}

	switch speedInput {
	case "s":
		fmt.Printf("%s✓ Slow speed selected!%s\n\n", types.ColorGreen, types.ColorReset)
		return 200, nil
	case "m":
		fmt.Printf("%s✓ Medium speed selected!%s\n\n", types.ColorGreen, types.ColorReset)
		return 100, nil
	case "f":
		fmt.Printf("%s✓ Fast speed selected!%s\n\n", types.ColorGreen, types.ColorReset)
		return 50, nil
	default:
		return 0, fmt.Errorf("invalid speed input: %s", speedInput)
	}
}

// displayControls shows the game controls and waits for user to start
func displayControls() {
	fmt.Printf("%s%sGame Controls:%s\n", types.ColorBold, types.ColorPurple, types.ColorReset)
	fmt.Printf("  %s↑%s Move Up:    %sW%s or %sK%s\n",
		types.ColorBlue, types.ColorReset, types.ColorYellow, types.ColorReset, types.ColorYellow, types.ColorReset)
	fmt.Printf("  %s↓%s Move Down:  %sS%s or %sJ%s\n",
		types.ColorBlue, types.ColorReset, types.ColorYellow, types.ColorReset, types.ColorYellow, types.ColorReset)
	fmt.Printf("  %s←%s Move Left:  %sA%s or %sH%s\n",
		types.ColorBlue, types.ColorReset, types.ColorYellow, types.ColorReset, types.ColorYellow, types.ColorReset)
	fmt.Printf("  %s→%s Move Right: %sD%s or %sL%s\n",
		types.ColorBlue, types.ColorReset, types.ColorYellow, types.ColorReset, types.ColorYellow, types.ColorReset)
	fmt.Printf("  %s🚪%s Quit Game:  %sQ%s or %sESC%s\n\n",
		types.ColorRed, types.ColorReset, types.ColorYellow, types.ColorReset, types.ColorYellow, types.ColorReset)

	fmt.Printf("%sPress %sENTER%s to start the game...%s", types.ColorWhite, types.ColorGreen, types.ColorWhite, types.ColorReset)
	var input string
	_, _ = fmt.Scanln(&input) // Explicitly ignore both return values to satisfy linters
}

func GetDifficultyInput() (types.Board, int, error) {
	// Initialize board
	var brd types.Board
	brd.Score = 0
	brd.Fruits = make(map[types.Point]struct{})

	// Display welcome screen
	displayWelcomeScreen()

	// Get board size
	width, height, err := getBoardSize()
	if err != nil {
		return types.Board{}, 0, err
	}
	brd.Width, brd.Height = width, height

	// Get game speed
	speed, err := getGameSpeed()
	if err != nil {
		return types.Board{}, 0, err
	}

	// Display controls and wait for user
	displayControls()

	return brd, speed, nil
}
