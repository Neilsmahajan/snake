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

			handleInput(char, key, inputChannel, s)
		}
	}
}

func handleInput(char rune, key keyboard.Key, inputChannel chan<- types.UserInput, s *types.Snake) {
	if key == keyboard.KeyEsc || char == 'q' || char == 'Q' {
		inputChannel <- types.UserInput{GamePlaying: false, Error: nil}
		return
	}

	var direction string
	switch char {
	case 'w', 'k':
		if s.Direction != types.DirectionDown {
			direction = types.DirectionUp
		}
	case 's', 'j':
		if s.Direction != types.DirectionUp {
			direction = types.DirectionDown
		}
	case 'a', 'h':
		if s.Direction != types.DirectionRight {
			direction = types.DirectionLeft
		}
	case 'd', 'l':
		if s.Direction != types.DirectionLeft {
			direction = types.DirectionRight
		}
	}

	if direction != "" {
		inputChannel <- types.UserInput{Direction: direction, GamePlaying: true, Error: nil}
	}
}
