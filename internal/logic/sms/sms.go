package sms

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"

	"goframe-websocket/internal/model"
	"goframe-websocket/internal/service"
)

type SSms struct {
}

func (S SSms) Check(ctx context.Context, channel string, key string, code string) bool {
	return S.GetCode(ctx, getSmsKey(channel, key)) == code
}

func (S SSms) Send(ctx context.Context, input model.SmsCodeInput) (*model.SmsSend, error) {
	res := &model.SmsSend{}
	key := getSmsKey(input.Channel, input.Mobile)
	filter, s := S.filter(input.Channel, input.Mobile)
	if filter && s != "" {
		return nil, gerror.New(s)
	}
	var code string
	if code = S.GetCode(ctx, key); code == "" {
		code = grand.Digits(6)
	}
	S.setCode(ctx, key, code)
	// ... 调取短信接口，建议异步任务执行 (暂无短信接口，省略处理)
	res.Code = code
	res.Type = input.Channel
	return res, nil

}

func (S SSms) GetCode(ctx context.Context, key string) string {
	str, _ := g.Redis().Do(ctx, "GET", key)
	return str.String()
}

func (S SSms) SetCode(ctx context.Context, key string, code string, exp uint) {
	//TODO implement me
	panic("implement me")
}

func (S SSms) DelCode(ctx context.Context, usage string, key string) {
	//TODO implement me
	panic("implement me")
}

func init() {
	sms := SSms{}
	service.RegisterSms(sms)
}

//获取短信key
func getSmsKey(channel string, mobile string) string {
	return fmt.Sprintf("sms_code:{%s}:{%s}", channel, mobile)
}

func (S SSms) filter(channel string, mobile string) (bool, string) {
	//过滤逻辑
	return false, ""
}

func (S SSms) setCode(ctx context.Context, key string, val string) {
	//考虑是否使用SETEX
	_, _ = g.Redis().Do(ctx, "SET", key, val)
}
