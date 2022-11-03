package websocket

import (
	"github.com/gogf/gf/v2/util/gconv"

	"goframe-websocket/internal/model"
	"goframe-websocket/internal/service"
)

func LoginController(client *Client, req *model.WsRequest) {
	UserId := gconv.Uint64(service.BizCtx().Get(client.context).User.Id)
	login := login{
		UserId: UserId,
		Client: client,
	}
	//TODO 2022-11-03
}
