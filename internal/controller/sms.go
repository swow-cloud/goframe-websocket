package controller

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "goframe-websocket/api/v1"
	"goframe-websocket/internal/consts"
	"goframe-websocket/internal/model"
	"goframe-websocket/internal/service"
)

var Sms = cSms{}

type cSms struct {
}

func (c *cSms) SmsCode(ctx context.Context, r *v1.SmsCodeReq) (res *v1.SmsCodeRes, err error) {

	channel := r.Channel
	mobile := r.Mobile
	switch channel {
	case consts.SmsLoginChannel:
	case consts.SmsForgetAccountChannel:
		existMobile, err := service.User().ExistMobile(ctx, mobile)
		if err != nil {
			return nil, err
		}
		if existMobile == false {
			return nil, gerror.New("账号不存在!")
		}
		break
	case consts.SmsRegister:
	case consts.SmsChangeAccountChannel:
		existMobile, err := service.User().ExistMobile(ctx, mobile)
		if err != nil {
			return nil, err
		}
		if existMobile {
			return nil, gerror.New("账号已被注册!")
		}
		break
	default:
		return nil, gerror.New("未知异常!")
	}
	res = &v1.SmsCodeRes{}
	send, err := service.Sms().Send(ctx, model.SmsCodeInput{
		Channel: channel,
		Mobile:  mobile,
	})
	if err != nil {
		return nil, err
	}
	res.Code = send.Code
	res.Debug = true
	return
}
