package service

import (
	"context"

	"goframe-websocket/internal/model"
	"goframe-websocket/internal/model/entity"
)

type IUser interface {
	Login(ctx context.Context, in model.UserLoginInput) (*model.UserToken, error)
	Logout(ctx context.Context) error
	GetUserByMobileAndPassword(ctx context.Context, mobile string, password string) (*entity.User, error)
	Register(ctx context.Context, in model.UserRegisterInput) (bool, error)
	CheckMobileUniq(ctx context.Context, mobile string) error
	GetUserInfo(ctx context.Context, id uint) (*entity.User, error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
