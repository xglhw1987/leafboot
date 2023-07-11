package msg

import (
	"github.com/name5566/leaf/network/json"
)

// var Processor network.Processor
var Processor = json.NewProcessor()

type RequestData struct {
}

type ResponseData struct {
}
type Hello struct {
	Name string
}

func init() {
	Processor.Register(&RequestData{})
	Processor.Register(&ResponseData{})
	Processor.Register(&Hello{})
}
