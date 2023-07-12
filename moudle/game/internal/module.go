package internal

import (
	"github.com/name5566/leaf/module"
	"leafboot/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	GameLogic.RegisteGameMsgCallback(skeleton)
}

func (m *Module) OnDestroy() {

}
