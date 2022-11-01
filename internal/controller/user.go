package controller

import (
	"context"

	v1 "goframe-websocket/api/v1"
	"goframe-websocket/internal/service"
)

var User = cUser{}

type cUser struct {
}

func (c *cUser) Info(ctx context.Context, req *v1.UserGetInfoReq) (res *v1.UserGetInfoRes, err error) {
	contextUser := service.BizCtx().Get(ctx).User
	res = &v1.UserGetInfoRes{}
	if err != nil {
		return nil, err
	}
	if contextUser != nil {
		res.Mobile = contextUser.Mobile
		res.Id = contextUser.Id
	}
	return res, nil
}
