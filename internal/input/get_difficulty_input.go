package input

import (
	"fmt"

	"github.com/neilsmahajan/snake/internal/board"
)

func GetDifficultyInput() (board.Board, int, error) {
	// Get the difficulty input from the user and return the board size and speed
	var speed int
	var brd board.Board
	fmt.Println("Welcome to the snake game written in Go!")

	fmt.Println("Please enter the board size ([s]mall, [m]edium, [l]arge):")
	var size string
	fmt.Scanln(&size)

	switch size {
	case "s":
		brd.Width, brd.Height = 20, 10
	case "m":
		brd.Width, brd.Height = 40, 20
	case "l":
		brd.Width, brd.Height = 80, 40
	default:
		return board.Board{}, 0, fmt.Errorf("invalid size input: %s", size)
	}
	fmt.Println("Please enter the speed ([s]low, [m]edium, [f]ast):")
	var speedInput string
	fmt.Scanln(&speedInput)
	switch speedInput {
	case "s":
		speed = 200
	case "m":
		speed = 100
	case "f":
		speed = 50
	default:
		return board.Board{}, 0, fmt.Errorf("invalid speed input: %s", speedInput)
	}
	fmt.Println("Use 'w' or 'k' to move up, 's' or 'j' to move down, 'a' or 'h' to move left, 'd' or 'l' to move right.")
	fmt.Println("Press 'q' or 'Esc' to quit the game.")
	return brd, speed, nil
}
