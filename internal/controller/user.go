package controller

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "goframe-websocket/api/v1"
	"goframe-websocket/internal/service"
)

var User = cUser{}

type cUser struct {
}

func (c *cUser) Info(ctx context.Context, req *v1.UserGetInfoReq) (res *v1.UserGetInfoRes, err error) {
	id := gconv.Int(service.Auth().GetIdentity(ctx))
	res = &v1.UserGetInfoRes{}
	userInfo, err := service.User().GetUserInfo(ctx, uint(id))
	if err != nil {
		return nil, err
	}
	if userInfo != nil {
		res.Mobile = userInfo.Mobile
		res.Id = userInfo.Id
	}
	return res, nil
}
