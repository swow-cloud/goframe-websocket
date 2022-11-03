package websocket

import (
	"fmt"
	"sync"

	"goframe-websocket/internal/model"
)

// ClientManager 客户端管理
type ClientManager struct {
	Clients         map[*Client]bool           //全部的连接
	ClientsLock     sync.RWMutex               //读写锁
	Users           map[string]*Client         //登录的用户 //uuid
	UserLock        sync.RWMutex               //读写锁
	Register        chan *Client               //连接处理
	Login           chan *login                //用户登录处理
	Unregister      chan *Client               //断开连接处理
	Broadcast       chan *model.WsResponse     //广播 向全部成员发送数据
	ClientBroadcast chan *model.WsResponse     //广播 向某个客户端发送数据
	TagBroadcast    chan *model.TagWsResponse  //广播 向某个标签成员发送数据
	UserBroadcast   chan *model.UserWsResponse //广播 向某个用户的所有连接发送数据
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients:       make(map[*Client]bool),
		Users:         make(map[string]*Client),
		Register:      make(chan *Client, 1000),
		Unregister:    make(chan *Client, 1000),
		Broadcast:     make(chan *model.WsResponse, 1000),
		TagBroadcast:  make(chan *model.TagWsResponse, 1000),
		UserBroadcast: make(chan *model.UserWsResponse, 1000),
	}
}

// GetUserKey 获取用户key
func GetUserKey(userId uint64) string {
	key := fmt.Sprintf("%s-%d", "ws", userId)
	return key
}

// InClient 客户端是否存在
func (manager *ClientManager) InClient(client *Client) (ok bool) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.Unlock()
	_, ok = manager.Clients[client]
	return
}
