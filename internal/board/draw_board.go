package board

import "fmt"

func DrawBoard(width, height, snakePositionX, snakePositionY int) {
	fmt.Print("\033[H\033[2J") // Clear the console
	for y := range height {
		for x := range width {
			if x == 0 || x == width-1 || y == 0 || y == height-1 {
				fmt.Print("#")
			} else if x == snakePositionX && y == snakePositionY {
				fmt.Print("0")
			} else {
				fmt.Print(" ")
			}
		}
		println()
	}
}
