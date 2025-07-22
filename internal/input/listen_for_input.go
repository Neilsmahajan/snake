package input

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/neilsmahajan/snake/internal/types"
)

const (
	directionUp    = "up"
	directionDown  = "down"
	directionLeft  = "left"
	directionRight = "right"
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
		if s.Direction != directionDown {
			direction = directionUp
		}
	case 's', 'j':
		if s.Direction != directionUp {
			direction = directionDown
		}
	case 'a', 'h':
		if s.Direction != directionRight {
			direction = directionLeft
		}
	case 'd', 'l':
		if s.Direction != directionLeft {
			direction = directionRight
		}
	}

	if direction != "" {
		inputChannel <- types.UserInput{Direction: direction, GamePlaying: true, Error: nil}
	}
}
