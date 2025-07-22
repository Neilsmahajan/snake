package input

import (
	"fmt"

	"github.com/neilsmahajan/snake/internal/types"
)

func GetDifficultyInput() (types.Board, int, error) {
	// Get the difficulty input from the user and return the board size and speed
	var speed int
	var brd types.Board
	brd.Score = 0
	brd.Fruits = make(map[types.Point]struct{})
	fmt.Println("Welcome to the snake game written in Go!")

	fmt.Println("Please enter the board size ([s]mall, [m]edium, [l]arge):")
	var size string
	if _, err := fmt.Scanln(&size); err != nil {
		return types.Board{}, 0, fmt.Errorf("error reading size input: %v", err)
	}

	switch size {
	case "s":
		brd.Width, brd.Height = 20, 10
	case "m":
		brd.Width, brd.Height = 40, 20
	case "l":
		brd.Width, brd.Height = 80, 40
	default:
		return types.Board{}, 0, fmt.Errorf("invalid size input: %s", size)
	}
	fmt.Println("Please enter the speed ([s]low, [m]edium, [f]ast):")
	var speedInput string
	if _, err := fmt.Scanln(&speedInput); err != nil {
		return types.Board{}, 0, fmt.Errorf("error reading speed input: %v", err)
	}
	switch speedInput {
	case "s":
		speed = 200
	case "m":
		speed = 100
	case "f":
		speed = 50
	default:
		return types.Board{}, 0, fmt.Errorf("invalid speed input: %s", speedInput)
	}
	fmt.Println("Use 'w' or 'k' to move up, 's' or 'j' to move down, 'a' or 'h' to move left, 'd' or 'l' to move right.")
	fmt.Println("Press 'q' or 'Esc' to quit the game.")
	return brd, speed, nil
}
