package service

import (
	"context"

	"github.com/yuwe1/pgim/internal/dao"
	"github.com/yuwe1/pgim/internal/model"
)

type appService struct{}

var AppService = new(appService)

// 获取设备信息
func (*appService) Get(ctx context.Context, appId int64) (*model.App, error) {

	// 从缓存中获取设备信息

	// 从数据库中获取app信息
	app, err := dao.AppDao.Get(appId)
	if err != nil {
		return app, nil
	}
	if app != nil {
		// 设置缓存
	}
	return app, nil

}
