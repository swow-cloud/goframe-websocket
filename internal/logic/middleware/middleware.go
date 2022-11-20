package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

	"goframe-websocket/internal/model"

	"goframe-websocket/internal/consts"
	"goframe-websocket/internal/service"
	"goframe-websocket/utility/response"
	"goframe-websocket/utility/utils"
)

type SMiddleware struct {
}

func (s *SMiddleware) Cors(request *ghttp.Request) {
	request.Response.CORSDefault()
	request.Middleware.Next()
}

func (s *SMiddleware) ResponseHandler(request *ghttp.Request) {
	request.Middleware.Next()
	if request.Response.BufferLength() > 0 {
		return
	}
	var (
		err             = request.GetError()
		res             = request.GetHandlerResponse()
		code gcode.Code = gcode.CodeOK
	)
	if err != nil {
		code = gerror.Code(err)
		code = gcode.New(consts.ApiFail, err.Error(), nil)
		response.JsonExit(request, code.Code(), code.Message(), utils.Charset.GetStack(err))

	} else {
		response.JsonExit(request, code.Code(), code.Message(), res)
	}
}

func (s *SMiddleware) Ctx(request *ghttp.Request) {
	//初始化
	customCtx := model.Context{
		Data: make(g.Map),
	}
	service.BizCtx().Init(request, &customCtx)
	request.Middleware.Next()
}

func (s *SMiddleware) Auth(request *ghttp.Request) {
	service.Auth().MiddlewareFunc()(request)
	ctx := request.GetCtx()

	id := gconv.Int(service.Auth().GetIdentity(ctx))
	userInfo, err := service.User().GetUserInfo(ctx, uint(id))

	if err != nil {
		g.Log().Error(ctx, "用户不存在!")
		request.Response.ClearBuffer()
		request.Response.WriteJsonExit(response.JsonRes{
			//todo 是否返回401
			Code:    401,
			Message: "用户不存在",
		})
	}
	//进行鉴权的时候设置用户上下文信息
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       userInfo.Id,
		Mobile:   userInfo.Mobile,
		Nickname: userInfo.Nickname,
		Avatar:   userInfo.Avatar,
	})

	request.Middleware.Next()
}

func init() {
	middleware := NewMiddleware()
	service.RegisterMiddleware(middleware)
}

func NewMiddleware() *SMiddleware {
	return &SMiddleware{}
}
