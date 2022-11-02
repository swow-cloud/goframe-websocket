package model

/**
用于处理ws响应
*/
import (
	"github.com/gogf/gf/v2/frame/g"
)

// WsRequest 当前输入对象
type WsRequest struct {
	Event string `json:"event"` //事件名称
	Data  g.Map  `json:"data"`  //数据
}

// WsResponse  输出对象
type WsResponse struct {
	Event string      `json:"event"` //事件名称
	Data  interface{} `json:"data"`  //数据
}

type TagWsResponse struct {
	Tag        string
	WsResponse *WsResponse
}

type UserWsResponse struct {
	UserID     uint64
	WsResponse *WsResponse
}

type ClientWsResponse struct {
	ID         string
	WsResponse *WsResponse
}
