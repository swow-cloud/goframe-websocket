// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint        `json:"id"        description:"用户ID"`
	Mobile    string      `json:"mobile"    description:"手机号"`
	Nickname  string      `json:"nickname"  description:"用户昵称"`
	Avatar    string      `json:"avatar"    description:"用户头像地址"`
	Gender    uint        `json:"gender"    description:"用户性别[0:未知;1:男;2:女;]"`
	Password  string      `json:"password"  description:"用户密码"`
	Motto     string      `json:"motto"     description:"用户座右铭"`
	Email     string      `json:"email"     description:"用户邮箱"`
	IsRobot   uint        `json:"isRobot"   description:"是否机器人[0:否;1:是;]"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
