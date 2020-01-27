package logic

import (
	"net"

	"github.com/yuwe1/pgim/pb"
	"github.com/yuwe1/pgim/rpcconf"
	"google.golang.org/grpc"
)

func StartRpcServer() {
	go func() {
		// recover恢复

		intListen, err := net.Listen("tcp", rpcconf.LogicConf.ClientRPCExtListenAddr)
		if err != nil {
			panic(err)
		}
		intServer := grpc.NewServer()
		pb.RegisterLogicClientExtServer(intServer, &LogicClientExt{})
		err = intServer.Serve(intListen)
		if err != nil {
			panic(err)
		}
	}()
}
