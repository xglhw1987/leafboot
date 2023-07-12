package gamelogic

import (
	"fmt"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"leafboot/base"
	"leafboot/msg"
)

func (f *FactoryGameLogic) RegisteGameMsgCallback(skeleton *module.Skeleton) {
	base.Skeleton = skeleton
	base.GameHandler(&msg.RequestData{}, f.handleRequestData)
}

func (f *FactoryGameLogic) handleRequestData(args []interface{}) {
	fmt.Printf("handleRequestData")
	log.Debug("handleRequestData")
}
func (f *FactoryGameLogic) handleResponseData(args []interface{}) {
	fmt.Printf("handleResponseData")
	log.Debug("handleResponseData")
}
