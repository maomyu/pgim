package basic

import (
	"github.com/yuwe1/pgim/pkg/client/rediscli/redispool"
	"github.com/yuwe1/pgim/pkg/client/dbpool"
	"github.com/yuwe1/pgim/pkg/config"
	"github.com/yuwe1/pgim/pkg/mq"
)

func Init() {
	config.Init()
	// rediscli.Init()
	redispool.Init()
	mq.Init()
	dbpool.Init()
}
