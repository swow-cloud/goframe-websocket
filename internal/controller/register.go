package controller

import (
	"context"

	v1 "goframe-websocket/api/v1"
	"goframe-websocket/internal/model"
	"goframe-websocket/internal/service"
)

var Register cRegister

type cRegister struct {
}

func (c *cRegister) Register(ctx context.Context, req *v1.RegisterDoReq) (res *v1.RegisterDoRes, err error) {
	res = &v1.RegisterDoRes{}
	//TODO 待验证sms check接口 2022-11-10
	_, err = service.User().Register(ctx, model.UserRegisterInput{
		Password: req.Password,
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		SmsCode:  req.SmsCode,
	})
	if err != nil {
		return nil, err
	}
	res.IsRegister = true
	return
}
