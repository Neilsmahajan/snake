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

func ListenForInput(inputChannel chan<- UserInput, direction *string) {
	if err := keyboard.Open(); err != nil {
		inputChannel <- UserInput{GamePlaying: false, Error: fmt.Errorf("error opening keyboard: %v", err)}
		return
	}
	defer keyboard.Close()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			inputChannel <- UserInput{GamePlaying: false, Error: fmt.Errorf("error reading input: %v", err)}
			return
		}

		if key == keyboard.KeyEsc || char == 'q' || char == 'Q' {
			inputChannel <- UserInput{GamePlaying: false, Error: nil}
			return
		}

		if (char == 'w' || char == 'k') && (*direction != "down") {
			inputChannel <- UserInput{Direction: "up", GamePlaying: true, Error: nil}
		} else if (char == 's' || char == 'j') && (*direction != "up") {
			inputChannel <- UserInput{Direction: "down", GamePlaying: true, Error: nil}
		} else if (char == 'a' || char == 'h') && (*direction != "right") {
			inputChannel <- UserInput{Direction: "left", GamePlaying: true, Error: nil}
		} else if (char == 'd' || char == 'l') && (*direction != "left") {
			inputChannel <- UserInput{Direction: "right", GamePlaying: true, Error: nil}
		}
	}
}
