package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leafboot/msg"
	"reflect"
)

func init() {
	handler(&msg.RequestData{}, handleRequest)
	handler(&msg.ResponseData{}, handleResponse)
	handler(&msg.Hello{}, handleHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleRequest(args []interface{}) {
	//log.Debug("server handlerRequest %+v", args)
	fmt.Printf("server handlerRequest %+v", args)
	m := args[0].(*msg.RequestData)
	a := args[1].(gate.Agent)
	fmt.Printf("handlerRequest m %+v a %+v", m, a)
	a.WriteMsg(m)
}

func handleResponse(args []interface{}) {
	log.Debug("server handleResponse %+v", args)
}
func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %v", m.Name)

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.Hello{
		Name: "client",
	})
}
