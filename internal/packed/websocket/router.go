package websocket

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"goframe-websocket/internal/model"
)

const (
	Error = "error"
	Login = "login"
	Join  = "join"
	Quit  = "quit"
	IsApp = "is_app"
	Ping  = "ping"
)

func Process(client *Client, message []byte) {
	defer func() {
		if r := recover(); r != nil {
			g.Log().Error(client.context, "处理数据Stop", r)
		}
	}()
	request := &model.WsRequest{}
	err := gconv.Struct(message, request)
	if err != nil {
		g.Log().Error(client.context, "数据解析失败 ", err)
	}
	//TODO 2022-11-02 待处理controller逻辑
	switch request.Event {
	case Login:
		break
	case Join:
		break
	case Quit:
		break
	case IsApp:
		break
	case Ping:
		break
	default:
		fmt.Println("Unknown Event")
		break
	}
}
