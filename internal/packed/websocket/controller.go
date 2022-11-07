package websocket

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"goframe-websocket/internal/model"
	"goframe-websocket/internal/service"
)

// LoginController  用户登录
func LoginController(client *Client, req *model.WsRequest) {
	UserId := gconv.Uint64(service.BizCtx().Get(client.context).User.Id)
	login := &login{
		UserId: UserId,
		Client: client,
	}
	clientManager.Login <- login
	client.SendMsg(&model.WsResponse{
		Event: Login,
		Data:  "success",
	})
}

func IsAppController(client *Client) {
	client.isApp = true
}

// JoinController 加入
func JoinController(client *Client, req *model.WsRequest) {
	name := gconv.String(req.Data["name"])
	if !client.tags.Contains(name) {
		client.tags.Append(name)
	}
	client.SendMsg(&model.WsResponse{
		Event: Join,
		Data:  client.tags.Slice(),
	})

}

func QuitController(client *Client, request *model.WsRequest) {
	name := gconv.String(request.Data["name"])
	if client.tags.Contains(name) {
		client.tags.RemoveValue(name)
	}
	client.SendMsg(&model.WsResponse{
		Event: Quit,
		Data:  client.tags.Slice(),
	})
}
func PingController(client *Client) {
	currentTime := uint64(gtime.Now().Unix())
	client.HeartBeat(currentTime)
}
