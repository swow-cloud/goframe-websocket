package user

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"goframe-websocket/internal/dao"
	"goframe-websocket/internal/model"
	"goframe-websocket/internal/model/entity"
	"goframe-websocket/internal/service"
	"goframe-websocket/utility/utils"
)

type SUser struct {
}

func (s *SUser) GetUserInfo(ctx context.Context, id uint) (*entity.User, error) {
	var user *entity.User
	err := dao.User.Ctx(ctx).Where("id", id).Scan(&user)
	if err != nil {
		return nil, gerror.New("用户不存在!")
	}
	return user, err
}

func (s *SUser) CheckMobileUniq(ctx context.Context, mobile string) error {
	var user *entity.User
	err := dao.User.Ctx(ctx).Where("mobile", mobile).Scan(&user)
	if err != nil {
		return err
	}
	if user != nil {
		return gerror.New("账号已被注册!")
	}
	return nil
}

func init() {
	user := NewSUser()
	service.RegisterUser(user)
}

func NewSUser() *SUser {
	return &SUser{}
}

func (s *SUser) Login(ctx context.Context, in model.UserLoginInput) (token *model.UserToken, err error) {
	userEntity, err := s.GetUserByMobileAndPassword(ctx, in.Mobile, in.Password)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, gerror.New(`手机号或密码错误`)
	}
	g.Log().Info(ctx, fmt.Sprintf("用户{%s}在【%s】登录成功", userEntity.Mobile, gtime.Datetime()))
	token = &model.UserToken{}
	token.Token, token.ExpireIn = service.Auth().LoginHandler(ctx)
	//设置上下文
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       userEntity.Id,
		Mobile:   userEntity.Mobile,
		Nickname: userEntity.Nickname,
		Avatar:   userEntity.Avatar,
	})
	return
}

func (s *SUser) Logout(ctx context.Context) error {
	service.Auth().LogoutHandler(ctx)
	return nil
}

// GetUserByMobileAndPassword 根据账号和密码查询用户信息，一般用于账号密码登录。
func (s *SUser) GetUserByMobileAndPassword(ctx context.Context, mobile string, password string) (*entity.User, error) {
	var user *entity.User
	get, _ := g.Cfg().Get(ctx, "app.secret")
	key := get.Val().(string)
	password, _ = utils.Encrypt([]byte(password), []byte(key))
	err := dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Mobile:   mobile,
		dao.User.Columns().Password: password,
	}).Scan(&user)
	return user, err
}

// Register 注册功能
func (s *SUser) Register(ctx context.Context, in model.UserRegisterInput) (bool, error) {
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var user *entity.User
		if err := gconv.Struct(in, &user); err != nil {
			return err
		}
		if err := s.CheckMobileUniq(ctx, user.Mobile); err != nil {
			return err
		}
		get, _ := g.Cfg().Get(ctx, "app.secret")
		key := get.Val().(string)
		encryptPassword, err := utils.Encrypt([]byte(user.Password), []byte(key))
		if err != nil {
			return err
		}
		user.Password = encryptPassword
		// user
		result, err := tx.Ctx(ctx).Insert("user", &user)
		if err != nil {
			g.Log().Error(ctx, fmt.Sprintf("用户{%s}在【%s】注册失败,失败原因【%s】", user.Mobile, gtime.Datetime(), err.Error()))
			return err
		}
		_, err = result.LastInsertId()
		if err != nil {
			return err
		}
		g.Log().Info(ctx, fmt.Sprintf("用户{%s}在【%s】注册成功", user.Mobile, gtime.Datetime()))
		return nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
