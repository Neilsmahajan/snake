package input

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func GetUserInput(direction string) (string, bool, error) {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		return direction, true, fmt.Errorf("Error reading input: %v", err)
	}
	if (char == 'w' || char == 'k') && direction != "down" {
		return "up", true, nil
	} else if (char == 's' || char == 'j') && direction != "up" {
		return "down", true, nil
	} else if (char == 'a' || char == 'h') && direction != "right" {
		return "left", true, nil
	} else if (char == 'd' || char == 'l') && direction != "left" {
		return "right", true, nil
	} else if char == 'q' || char == 'Q' {
		return direction, false, nil // Return the current direction and indicate game over
	}
	return direction, true, nil // Return the current direction and indicate game continues
}
