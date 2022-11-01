package model

import "github.com/gogf/gf/v2/frame/g"

const (
	ContextKey = "ContextKey"
)

type Context struct {
	User *ContextUser // 上下文用户信息
	Data g.Map        // 自定KV变量，业务模块根据需要设置，不固定
}

type ContextUser struct {
	Id       uint
	Mobile   string
	Nickname string
	Avatar   string
}
