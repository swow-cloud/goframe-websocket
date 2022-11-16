package websocket

import (
	"context"
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"

	"goframe-websocket/internal/model"
)

// ClientManager 客户端管理
type ClientManager struct {
	Clients         map[*Client]bool             //全部的连接
	ClientsLock     sync.RWMutex                 //读写锁
	Users           map[string]*Client           //登录的用户 //uuid
	UserLock        sync.RWMutex                 //读写锁
	Register        chan *Client                 //连接处理
	Login           chan *login                  //用户登录处理
	Unregister      chan *Client                 //断开连接处理
	Broadcast       chan *model.WsResponse       //广播 向全部成员发送数据
	ClientBroadcast chan *model.ClientWsResponse //广播 向某个客户端发送数据
	TagBroadcast    chan *model.TagWsResponse    //广播 向某个标签成员发送数据
	UserBroadcast   chan *model.UserWsResponse   //广播 向某个用户的所有连接发送数据
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients:       make(map[*Client]bool),
		Users:         make(map[string]*Client),
		Register:      make(chan *Client, 1000),
		Unregister:    make(chan *Client, 1000),
		Login:         make(chan *login, 1000),
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
	defer manager.ClientsLock.RUnlock()
	_, ok = manager.Clients[client]
	return
}

// GetClients 获取所有客户端
func (manager *ClientManager) GetClients() (clients map[*Client]bool) {
	clients = make(map[*Client]bool)
	manager.ClientsRange(func(client *Client, value bool) (result bool) {
		clients[client] = value
		return true
	})
	return
}

// ClientsRange 遍历
func (manager *ClientManager) ClientsRange(f func(client *Client, value bool) (result bool)) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	for key, value := range manager.Clients {
		result := f(key, value)
		if result == false {
			return
		}
	}
	return
}

// GetClientsLen 获取连接总数
func (manager *ClientManager) GetClientsLen() (clientsLen int) {
	return len(manager.Clients)
}

// AddClients 添加客户端
func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()

	manager.Clients[client] = true
}

// DelClients 删除客户端
func (manager *ClientManager) DelClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	if _, ok := manager.Clients[client]; ok {
		delete(manager.Clients, client)
	}
}

// GetUserClient 获取用户的连接
func (manager *ClientManager) GetUserClient(userId uint64) (client *Client) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	userKey := GetUserKey(userId)
	if c, ok := manager.Users[userKey]; ok {
		client = c
	}
	return
}

// AddUsers 添加用户客户端
func (manager *ClientManager) AddUsers(key string, client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	manager.Users[key] = client
}

// DelUsers 删除用户客户端
func (manager *ClientManager) DelUsers(client *Client) (result bool) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	key := GetUserKey(client.UserId)
	if value, ok := manager.Users[key]; ok {
		if value.Addr != client.Addr {
			return
		}
		delete(manager.Users, key)
		result = true
	}
	return
}

// GetUsersLen 获取客户端连结数
func (manager *ClientManager) GetUsersLen() (userLen int) {
	return len(manager.Users)
}

// EventRegister 用户建立连接事件
func (manager *ClientManager) EventRegister(client *Client) {
	//UserId := gconv.Uint64(service.BizCtx().Get(client.context).User.Id)
	manager.AddClients(client)
	//发送当前客户端ID
	client.SendMsg(&model.WsResponse{
		Event: "connect",
		Content: g.Map{
			"ping_interval": 20,
			"ping_timeout":  60,
		},
	})
}

// EventLogin 用户登录
func (manager *ClientManager) EventLogin(login *login) {
	client := login.Client
	if manager.InClient(client) {
		userKey := login.GetKey()
		manager.AddUsers(userKey, client)
	}
}

// EventUnregister 断开连接事件
func (manager *ClientManager) EventUnregister(client *Client) {
	manager.DelClients(client)
	//删除用户连接
	result := manager.DelUsers(client)

	if result == false {
		//不是当前连接客户端
		return
	}
	//关闭 chan
	close(client.Send)
}

// ClearTimeoutConnections 定时清理超时连接
func (manager *ClientManager) ClearTimeoutConnections() {
	currentTime := uint64(gtime.Now().Unix())
	clients := clientManager.GetClients()
	for client := range clients {
		if client.IsHeartbeatTimeout(currentTime) {
			//断开连接
			_ = client.Socket.Close()
		}
	}
}

func (manager *ClientManager) ping(ctx context.Context) {
	//通过定时任务，发送心跳
	_, _ = gcron.Add(ctx, "0 */1 * * * *", func(ctx context.Context) {
		res := &model.WsResponse{
			Event:   Ping,
			Content: g.Map{},
		}
		sendToAll(res)
	})

	//定时任务，清理超时连接
	_, _ = gcron.Add(ctx, "0 */30 * * * *", func(ctx context.Context) {
		manager.ClearTimeoutConnections()
	})

}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.Register:
			//用户建立连接
			manager.EventRegister(conn)
			break
		case login := <-manager.Login:
			//用户登录事件
			manager.EventLogin(login)
			break
		case conn := <-manager.Unregister:
			//断开连接时间
			manager.EventUnregister(conn)
			break
		case message := <-manager.Broadcast:
			//全部客户端广播事件
			clients := manager.GetClients()
			for client := range clients {
				client.SendMsg(message)
			}
			break
		case message := <-manager.TagBroadcast:
			//标签广播事件
			clients := manager.GetClients()
			for client := range clients {
				client.SendMsg(message.WsResponse)
			}
			break
		case message := <-manager.UserBroadcast:
			//用户广播事件
			clients := manager.GetClients()
			for client := range clients {
				if client.UserId == message.UserID {
					client.SendMsg(message.WsResponse)
				}
			}
			break
		case message := <-manager.ClientBroadcast:
			clients := manager.GetClients()
			// 单个客户端广播事件
			for client := range clients {
				if client.ID == message.ID {
					client.SendMsg(message.WsResponse)
				}
			}
			break
		}
	}
}

//sendToAll  发送全部客户端
func sendToAll(response *model.WsResponse) {
	clientManager.Broadcast <- response
}

// SendToClientID 发送到指定客户端
func SendToClientID(id string, response *model.WsResponse) {
	clientRes := &model.ClientWsResponse{
		ID:         id,
		WsResponse: response,
	}
	clientManager.ClientBroadcast <- clientRes
}

// SendToUser 发送到指定用户
func SendToUser(userID uint64, response *model.WsResponse) {
	userRes := &model.UserWsResponse{
		UserID:     userID,
		WsResponse: response,
	}
	clientManager.UserBroadcast <- userRes
}

// SendToTag 发送标签
func SendToTag(tag string, response *model.WsResponse) {
	tagRes := &model.TagWsResponse{
		Tag:        tag,
		WsResponse: response,
	}
	clientManager.TagBroadcast <- tagRes
}
