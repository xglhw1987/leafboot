package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"leafboot/appmain/gamelogic"
	"leafboot/base"
	"leafboot/conf"
	"leafboot/moudle/game"
	"leafboot/moudle/gate"
	"leafboot/moudle/login"
	"leafboot/moudle/robot"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath
	base.GameChanRPC = game.ChanRPC
	base.RobotChanRPC = robot.ChanRPC

	gamelogic.StartGame(&gamelogic.FactoryGameLogic{})
	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
