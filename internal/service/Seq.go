package service

import (
	"context"

	"github.com/yuwe1/pgim/internal/cache"
)

type seqService struct{}

var SeqService = new(seqService)

// 获取下一个序列号
func (*seqService) GetUserNext(ctx context.Context, appId, userId int64) (int64, error) {
	return cache.SeqCache.Incr(cache.SeqCache.UserSeqKey(appId, userId))
}
