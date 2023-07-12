package base

import (
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/module"
	"reflect"
)

var (
	Skeleton     *module.Skeleton
	GameChanRPC  *chanrpc.Server //其他模块，通过GameChanRPC给主逻辑发消息
	RobotChanRPC *chanrpc.Server //其他模块，通过RobotChanRPC给机器人模块发消息
)

// 抽象游戏逻辑类
type IGameLogic interface {
	InitAppMain(appMain IGameLogic)
	InitBoot(appMain IGameLogic)
	Start(netgate *gate.Gate) error
	Gamestart(table ITable)
	OnCreateTable() ITable
	RegisteGameMsgCallback(skeleton *module.Skeleton)
}

type IRobot interface {
	InitAppMain(appMain IRobot)
	HandleRobotMsg(args []interface{})
	OnCreateRobot() IPlayerNode
	RegisteRobotMsg()
	OnRobotLoginIn(player IPlayerNode, loginmsg interface{})
	OnRobotLoginOut(player IPlayerNode)
	CloseRobot(player IPlayerNode)
	OnDestroy()
}

// 抽象个人类
type IPlayerNode interface {
}

// 抽象桌子类
type ITable interface {
}

func GameHandler(m interface{}, h interface{}) {
	Skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
