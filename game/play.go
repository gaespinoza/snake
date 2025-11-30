package game

import "github.com/gaespinoza/snake/models"

func TakeStep(game *models.Game) error {
	return game.MoveSnake()
}
