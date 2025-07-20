package snake

func MoveSnake(snakePositionX, snakePositionY, width, height int, direction string) (int, int, bool) {
	switch direction {
	case "up":
		snakePositionY--
	case "down":
		snakePositionY++
	case "left":
		snakePositionX--
	case "right":
		snakePositionX++
	case "still":
		// Do nothing, snake stays in the same position
	}
	if snakePositionX <= 0 || snakePositionX >= width-1 || snakePositionY <= 0 || snakePositionY >= height-1 {
		return snakePositionX, snakePositionY, false // Game over if the snake hits the wall
	}
	return snakePositionX, snakePositionY, true
}
