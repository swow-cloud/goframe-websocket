package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterDoReq struct {
	g.Meta   `path:"/register" method:"post"`
	Mobile   string `json:"mobile" v:"required#请输入手机号"`
	Password string `json:"password" v:"required#请输入密码" `
	Nickname string `json:"nickname" v:"required#请输入昵称"`
	SmsCode  string `json:"smscode"  v:"required#请输入验证码"`
}

type RegisterDoRes struct {
	IsRegister bool `json:"isRegister"`
}
