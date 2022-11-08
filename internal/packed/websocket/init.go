package websocket

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/websocket"
)

var (
	clientManager = NewClientManager()
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartWebsocket(ctx context.Context) {
	g.Log().Info(ctx, "启动: WebSocket")
	go clientManager.start()
	go clientManager.ping(ctx)

}

func WsHandler(request *ghttp.Request) {
	conn, err := upGrader.Upgrade(request.Response.ResponseWriter, request.Request, nil)
	if err != nil {
		return
	}
	currentTime := uint64(gtime.Now().Unix())
	client := NewClient(request.GetCtx(), conn.RemoteAddr().String(), conn, currentTime)
	go client.read()
	go client.write()
	//用户连接事件
	clientManager.Register <- client

}
