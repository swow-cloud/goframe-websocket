package service

import (
	"context"

	"goframe-websocket/internal/model"
)

type ISms interface {
	Send(ctx context.Context, input model.SmsCodeInput) (*model.SmsSend, error)
	GetCode(ctx context.Context, key string) string
	SetCode(ctx context.Context, key string, code string, exp uint)
	DelCode(ctx context.Context, channel string, key string)
	Check(ctx context.Context, channel string, key string, code string) bool
}

var localSms ISms

func Sms() ISms {
	if localSms == nil {
		panic("implement not found for interface localSMs, forgot register?")
	}
	return localSms
}

func RegisterSms(i ISms) {
	localSms = i
}
