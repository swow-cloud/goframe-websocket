package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"

	"goframe-websocket/internal/controller"
	"goframe-websocket/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.Hello,
					controller.Login,
					controller.Register,
				)
				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.ALLMap(g.Map{
						"/user/info": controller.User.Info,
					})
					group.ALL("/ws", 1)
				})
			})
			//s.BindHandler("/ws", func(r *ghttp.Request) {
			//	var ctx = r.Context()
			//	ws, err := r.WebSocket()
			//	if err != nil {
			//		glog.Error(ctx, err)
			//		r.Exit()
			//	}
			//	for {
			//		msgType, msg, err := ws.ReadMessage()
			//		if err != nil {
			//			return
			//		}
			//		if err = ws.WriteMessage(msgType, msg); err != nil {
			//			return
			//		}
			//	}
			//})
			s.SetServerRoot(gfile.MainPkgPath())
			s.SetPort(8199)
			s.Run()
			return nil
		},
	}
)
