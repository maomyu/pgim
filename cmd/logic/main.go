package main

import (
	"github.com/yuwe1/pgim/api/logic"
	"github.com/yuwe1/pgim/pkg/client/dbpool"
	"github.com/yuwe1/pgim/pkg/logger"
	"github.com/yuwe1/pgim/pkg/util"

	common "github.com/yuwe1/pgim/pkg"
)

func main() {
	common.Init()
	session, err, p, c := dbpool.GetSession()
	defer func() {
		if session != nil {
			session.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	db := session.DB
	// 初始化自增id配置
	util.InitUID(db)
	// 初始化rpc client

	// 初始化rpc服务
	logic.StartRpcServer()
	logger.Logger.Info("logic server start")
	select {}
}
