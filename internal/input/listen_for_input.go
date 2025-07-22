package input

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/neilsmahajan/snake/internal/types"
)

func ListenForInput(inputChannel chan<- types.UserInput, s *types.Snake, stopChannel <-chan struct{}) {
	if err := keyboard.Open(); err != nil {
		inputChannel <- types.UserInput{GamePlaying: false, Error: fmt.Errorf("error opening keyboard: %v", err)}
		return
	}
	defer keyboard.Close()

	for {
		select {
		case <-stopChannel:
			return
		default:
			char, key, err := keyboard.GetKey()
			if err != nil {
				inputChannel <- types.UserInput{GamePlaying: false, Error: fmt.Errorf("error reading input: %v", err)}
				return
			}

			if key == keyboard.KeyEsc || char == 'q' || char == 'Q' {
				inputChannel <- types.UserInput{GamePlaying: false, Error: nil}
				return
			}

			if (char == 'w' || char == 'k') && (s.Direction != "down") {
				inputChannel <- types.UserInput{Direction: "up", GamePlaying: true, Error: nil}
			} else if (char == 's' || char == 'j') && (s.Direction != "up") {
				inputChannel <- types.UserInput{Direction: "down", GamePlaying: true, Error: nil}
			} else if (char == 'a' || char == 'h') && (s.Direction != "right") {
				inputChannel <- types.UserInput{Direction: "left", GamePlaying: true, Error: nil}
			} else if (char == 'd' || char == 'l') && (s.Direction != "left") {
				inputChannel <- types.UserInput{Direction: "right", GamePlaying: true, Error: nil}
			}
		}
	}
}
