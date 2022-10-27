package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"goframe-websocket/internal/cmd"
	_ "goframe-websocket/internal/logic"
	_ "goframe-websocket/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
}
