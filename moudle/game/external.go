package game

import (
	"leafboot/base"
	"leafboot/moudle/game/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

func InitLogic(logic base.IGameLogic) {
	internal.GameLogic = logic
	internal.GameLogic.InitAppMain(logic)
}
