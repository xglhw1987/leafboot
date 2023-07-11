package gate

import (
	"leafboot/game"
	"leafboot/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.RequestData{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.ResponseData{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}
