package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"goframe-websocket/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("哈喽个der啊")
	return
}
