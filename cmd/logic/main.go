package main

import (
	"github.com/yuwe1/pgim/api/logic"
	"github.com/yuwe1/pgim/pkg/logger"
)

func main() {
	pkg.Init()

	// 初始化rpc client

	// 初始化rpc服务
	logic.StartRpcServer()
	logger.Logger.Info("logic server start")
	select {}
}
