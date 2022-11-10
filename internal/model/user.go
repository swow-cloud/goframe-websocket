package model

import (
	"time"
)

// UserLoginInput 用户登录
type UserLoginInput struct {
	Mobile   string //账号
	Password string //密码
}

// UserRegisterInput 用户注册
type UserRegisterInput struct {
	Mobile   string
	Password string
	SmsCode  string
	Nickname string
}

// SmsCodeInput 短信验证码
type SmsCodeInput struct {
	Mobile  string
	Channel string
}

type SmsSend struct {
	Type string
	Code string
}

type UserToken struct {
	Token    string
	ExpireIn time.Time
}
