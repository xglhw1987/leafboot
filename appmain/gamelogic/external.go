package gamelogic

import (
	"github.com/name5566/leaf/gate"
	"leafboot/base"
)

type FactoryGameLogic struct {
	AppLogic base.IGameLogic
}

func (f *FactoryGameLogic) InitAppMain(appMain base.IGameLogic) {
	f.AppLogic = appMain
}
func (f *FactoryGameLogic) InitBoot(appMain base.IGameLogic) {

}
func (f *FactoryGameLogic) Start(netgate *gate.Gate) error {
	return nil
}
func (f *FactoryGameLogic) Gamestart(table base.ITable) {

}
func (f *FactoryGameLogic) OnCreateTable() base.ITable {
	return nil
}
