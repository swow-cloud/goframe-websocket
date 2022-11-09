package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" method:"post"`
	Mobile   string `json:"mobile" v:"required#请输入手机号"`
	Password string `json:"password" v:"required#请输入密码"`
}

type LoginDoRes struct {
	Token    string `json:"token"`
	ExpireIn string `json:"expireIn"`
}
