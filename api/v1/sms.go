package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SmsCodeReq struct {
	g.Meta  `path:"/sms/code" method:"post" tag:"内容"`
	Channel string `json:"channel" v:"required|in:login,register,forget_account,change_account#请输入通道|通道错误"`
	Mobile  string `json:"mobile" v:"required#请输入手机号"`
}

type SmsCodeRes struct {
	Code  string `json:"code"`
	Debug bool   `json:"debug"`
}
