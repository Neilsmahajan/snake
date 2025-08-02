package input

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/neilsmahajan/snake/internal/types"
)

func ListenForInput(inputChannel chan<- types.UserInput, s *types.Snake, stopChannel <-chan struct{}) {
	if err := keyboard.Open(); err != nil {
		inputChannel <- types.UserInput{GamePlaying: false, Error: fmt.Errorf("error opening keyboard: %v", err)}
		return
	}
	defer func() {
		_ = keyboard.Close() // #nosec G104 - keyboard cleanup, error doesn't affect game flow
		// Small delay to ensure keyboard cleanup completes
		time.Sleep(10 * time.Millisecond)
	}()

	for {
		select {
		case <-stopChannel:
			return
		default:
			char, key, err := keyboard.GetKey()
			if err != nil {
				// Don't send error if we're stopping (channel might be closed)
				select {
				case <-stopChannel:
					return
				default:
					inputChannel <- types.UserInput{GamePlaying: false, Error: fmt.Errorf("error reading input: %v", err)}
					return
				}
			}

			// Check if we should stop before sending input
			select {
			case <-stopChannel:
				return
			default:
				handleInput(char, key, inputChannel, s)
			}
		}
	}
}

func handleInput(char rune, key keyboard.Key, inputChannel chan<- types.UserInput, s *types.Snake) {
	if key == keyboard.KeyEsc || char == 'q' || char == 'Q' {
		// Use non-blocking send to avoid hanging if channel is full/closed
		select {
		case inputChannel <- types.UserInput{GamePlaying: false, Error: nil}:
		default:
			// Channel is full or closed, just return
		}
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
		// Use non-blocking send to avoid hanging if channel is full/closed
		select {
		case inputChannel <- types.UserInput{Direction: direction, GamePlaying: true, Error: nil}:
		default:
			// Channel is full or closed, just return
		}
	}
}
