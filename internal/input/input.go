package input

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

type UserInput struct {
	Direction   string
	GamePlaying bool
	Error       error
}

func GetUserInput(direction string, inputChannel chan<- UserInput) {
	char, key, err := keyboard.GetSingleKey()
	if err != nil {
		inputChannel <- UserInput{Direction: direction, GamePlaying: false, Error: fmt.Errorf("Error reading input: %v", err)}
	} else if key == keyboard.KeyEsc || char == 'q' || char == 'Q' {
		inputChannel <- UserInput{Direction: direction, GamePlaying: false, Error: nil}
	} else if (char == 'w' || char == 'k') && direction != "down" {
		inputChannel <- UserInput{Direction: "up", GamePlaying: true, Error: nil}
	} else if (char == 's' || char == 'j') && direction != "up" {
		inputChannel <- UserInput{Direction: "down", GamePlaying: true, Error: nil}
	} else if (char == 'a' || char == 'h') && direction != "right" {
		inputChannel <- UserInput{Direction: "left", GamePlaying: true, Error: nil}
	} else if (char == 'd' || char == 'l') && direction != "left" {
		inputChannel <- UserInput{Direction: "right", GamePlaying: true, Error: nil}
	} else {
		inputChannel <- UserInput{Direction: direction, GamePlaying: true, Error: nil}
	}
}
