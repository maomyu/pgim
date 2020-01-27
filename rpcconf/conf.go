package rpcconf

// logic配置
var (
	LogicConf logicConf
	ConnConf  connConf
	WSConf    wsConf
)

type logicConf struct {
	// 监听的地址
	RPCIntListenAddr       string
	ClientRPCExtListenAddr string
	ServerRPCExtListenAddr string
	ConnRPCAddrs           string
}

// conn配置
type connConf struct {
	TCPListenAddr string
	RPCListenAddr string
	LocalAddr     string
	LogicRPCAddrs string
}

// ws
type wsConf struct {
	WSListenAddr  string
	RPCListenAddr string
	LocalAddr     string
	LogicRPCAddrs string
}

func init() {
	LogicConf = logicConf{
		RPCIntListenAddr:       ":50000",
		ClientRPCExtListenAddr: ":50001",
		ServerRPCExtListenAddr: ":50002",
		ConnRPCAddrs:           "127.0.0.1:60000,127.0.0.1:60001",
	}

	ConnConf = connConf{
		TCPListenAddr: ":8080",
		RPCListenAddr: ":60000",
		LocalAddr:     "127.0.0.1:60000",
		LogicRPCAddrs: "127.0.0.1:50000",
	}

	WSConf = wsConf{
		WSListenAddr:  ":8081",
		RPCListenAddr: ":60001",
		LocalAddr:     "127.0.0.1:60001",
		LogicRPCAddrs: "127.0.0.1:50000",
	}
}
