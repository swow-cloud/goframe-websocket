package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserGetInfoReq struct {
	g.Meta `path:"/info" method:"get"`
}

type UserGetInfoRes struct {
	Id     uint   `json:"id"`
	Mobile string `json:"mobile"`
}
