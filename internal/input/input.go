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

func ListenForInput(inputChannel chan<- UserInput) {
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

		switch char {
		case 'w', 'k':
			inputChannel <- UserInput{Direction: "up", GamePlaying: true, Error: nil}
		case 's', 'j':
			inputChannel <- UserInput{Direction: "down", GamePlaying: true, Error: nil}
		case 'a', 'h':
			inputChannel <- UserInput{Direction: "left", GamePlaying: true, Error: nil}
		case 'd', 'l':
			inputChannel <- UserInput{Direction: "right", GamePlaying: true, Error: nil}
		}
	}
}
