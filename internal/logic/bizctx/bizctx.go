package bizctx

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"goframe-websocket/internal/model"
	"goframe-websocket/internal/service"
)

type SBizCtx struct {
}

func (s *SBizCtx) SetData(ctx context.Context, data g.Map) {
	s.Get(ctx).Data = data
}

func init() {
	service.RegisterBizCtx(New())
}

func New() *SBizCtx {
	return &SBizCtx{}
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *SBizCtx) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(model.ContextKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *SBizCtx) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *SBizCtx) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}
