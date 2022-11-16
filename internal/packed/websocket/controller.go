package websocket

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-websocket/internal/service"

	"goframe-websocket/internal/model"
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
		Event:   Login,
		Content: "success",
	})
}

func IsAppController(client *Client) {
	client.isApp = true
}

// JoinController 加入
func JoinController(client *Client, req *model.WsRequest) {
	name := gconv.String(req.Content["name"])
	if !client.tags.Contains(name) {
		client.tags.Append(name)
	}
	client.SendMsg(&model.WsResponse{
		Event:   Join,
		Content: client.tags.Slice(),
	})

}

func QuitController(client *Client, request *model.WsRequest) {
	name := gconv.String(request.Content["name"])
	if client.tags.Contains(name) {
		client.tags.RemoveValue(name)
	}
	client.SendMsg(&model.WsResponse{
		Event:   Quit,
		Content: client.tags.Slice(),
	})
}
func PingController(client *Client) {
	currentTime := uint64(gtime.Now().Unix())
	client.HeartBeat(currentTime)
}
