package input

import "fmt"

func GetUserInput(direction string) (string, bool) {
	var input string
	fmt.Scanln(&input)
	if (input == "w" || input == "k") && direction != "down" {
		return "up", false
	} else if (input == "s" || input == "j") && direction != "up" {
		return "down", false
	} else if (input == "a" || input == "h") && direction != "right" {
		return "left", false
	} else if (input == "d" || input == "l") && direction != "left" {
		return "right", false
	} else if input == "q" || input == "Q" {
		return direction, true // Return the current direction and indicate game over
	}
	return direction, false // Return the current direction and indicate game continues
}
