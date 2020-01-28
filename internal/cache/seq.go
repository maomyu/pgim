package cache

import (
	"strconv"

	"github.com/garyburd/redigo/redis"
	"github.com/yuwe1/pgim/pkg/client/rediscli/redispool"
)

const (
	UserSeqKey  = "user_seq"
	GroupSeqKey = "group_seq"
)

type seqCache struct {
}

var SeqCache = new(seqCache)

func (*seqCache) UserSeqKey(appId, userId int64) string {
	return UserSeqKey + strconv.FormatInt(appId, 10) + ":" + strconv.FormatInt(userId, 10)
}
func (*seqCache) GroupKey(appId, groupId int64) string {
	return GroupSeqKey + strconv.FormatInt(appId, 10) + ":" + strconv.FormatInt(groupId, 10)
}

// 将序列号进行自增
func (c *seqCache) Incr(key string) (int64, error) {
	f, _, p, co := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, co)
		}
	}()
	conn := f.GetConn()
	if result, _ := redis.Int64(conn.Do("INCR", key)); result > 0 {
		return result, nil
	}
	return 0, nil
}
