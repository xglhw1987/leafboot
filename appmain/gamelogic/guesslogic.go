package gamelogic

import (
	"leafboot/base"
	"leafboot/moudle/game"
)

func StartGame(flogic base.IGameLogic) {
	game.InitLogic(flogic)
}
