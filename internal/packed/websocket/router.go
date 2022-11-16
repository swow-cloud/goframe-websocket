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
	Ping  = "heartbeat"
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
		return
	}
	switch request.Event {
	case Login:
		LoginController(client, request)
		break
	case Join:
		JoinController(client, request)
		break
	case Quit:
		QuitController(client, request)
		break
	case IsApp:
		IsAppController(client)
		break
	case Ping:
		PingController(client)
		break
	default:
		fmt.Println("Unknown Event")
		break
	}
}
