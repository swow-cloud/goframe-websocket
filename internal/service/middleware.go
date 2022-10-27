package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type IMiddleware interface {
	ResponseHandler(request *ghttp.Request)
	Ctx(request *ghttp.Request)
	Auth(request *ghttp.Request)
	Cors(request *ghttp.Request)
}

var localMiddleware IMiddleware

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(m IMiddleware) {
	localMiddleware = m
}
