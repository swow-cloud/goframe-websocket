package websocket

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/gorilla/websocket"

	"goframe-websocket/internal/model"
)

const (
	//心跳超时时间
	heartbeatExpirationTime int = 6 * 60
)

type login struct {
	UserId uint64
	Client *Client
}

type Client struct {
	Addr          string                 //客户端地址
	ID            string                 //连接唯一标识
	Socket        *websocket.Conn        //用户连接对象
	Send          chan *model.WsResponse //待发送的数据
	SendCLose     bool                   //发送是否关闭
	UserId        uint64                 //用户ID,用户登录才有
	FirstTime     uint64                 //首次连接时间
	HeartbeatTime uint64                 //用户上次心跳时间
	LoginTime     uint64                 //登录时间
	IsApp         bool                   //是否是APP
	tags          garray.StrArray        //标签
	context       context.Context        //自定义上下文
}

func NewClient(ctx context.Context, addr string, socket *websocket.Conn, firstTime uint64) (client *Client) {
	return &Client{
		context:       ctx,
		Addr:          addr,
		ID:            guid.S(),
		Socket:        socket,
		Send:          make(chan *model.WsResponse, 100),
		SendCLose:     false,
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
	}
}

func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			g.Log().Error(c.context, "write stop", string(debug.Stack()), r)
		}
	}()
	defer func() {
		c.close()
	}()
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(message)
		//TODO 2022-11-02
		Process(c, message)
	}
}

func (c *Client) close() {
	if c.SendCLose {
		return
	}
	c.SendCLose = true
	close(c.Send)
}
