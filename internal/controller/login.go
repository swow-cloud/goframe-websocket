package controller

import (
	"context"

	v1 "goframe-websocket/api/v1"
	"goframe-websocket/internal/model"
	"goframe-websocket/internal/service"
)

var (
	Login cLogin
)

type cLogin struct {
}

func (c *cLogin) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginDoRes, err error) {
	res = &v1.LoginDoRes{}
	userToken, err := service.User().Login(ctx, model.UserLoginInput{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	res.Token, res.ExpireIn = userToken.Token, userToken.ExpireIn
	return

}
