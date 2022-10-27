package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"

	"goframe-websocket/internal/consts"
	"goframe-websocket/internal/service"
	"goframe-websocket/utility/response"
	"goframe-websocket/utility/utils"
)

type sMiddleware struct {
}

func (s *sMiddleware) Cors(request *ghttp.Request) {
	request.Response.CORSDefault()
	request.Middleware.Next()
}

func (s *sMiddleware) ResponseHandler(request *ghttp.Request) {
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

func (s *sMiddleware) Ctx(request *ghttp.Request) {
	request.Middleware.Next()
}

func (s *sMiddleware) Auth(request *ghttp.Request) {
	service.Auth().MiddlewareFunc()(request)
	request.Middleware.Next()
}

func init() {
	middleware := NewMiddleware()
	service.RegisterMiddleware(middleware)
}

func NewMiddleware() *sMiddleware {
	return &sMiddleware{}
}
