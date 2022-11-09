package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"

	"goframe-websocket/internal/controller"
	"goframe-websocket/internal/packed/websocket"
	"goframe-websocket/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			//启动服务
			websocket.StartWebsocket(ctx)

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
					group.ALL("/ws", websocket.WsHandler)
				})

			})
			s.SetServerRoot(gfile.MainPkgPath())
			s.SetPort(8199)
			s.Run()
			return nil
		},
	}
)
