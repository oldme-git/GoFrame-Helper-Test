package main

import (
	_ "GoFrame-Helper-Test/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"GoFrame-Helper-Test/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
