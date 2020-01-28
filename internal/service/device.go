package service

import (
	"context"
	"fmt"

	"github.com/yuwe1/pgim/internal/dao"
	"github.com/yuwe1/pgim/internal/model"
	"github.com/yuwe1/pgim/pkg/logger"

	"github.com/yuwe1/pgim/pkg/util"
)

type deviceService struct{}

var DeviceService = new(deviceService)

// 注册设备
func (*deviceService) Register(ctx context.Context, device model.Device) (int64, error) {

	// 获取app信息，判断接入app是否可信
	app, err := AppService.Get(ctx, device.AppId)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	if app == nil {
		return 0, fmt.Errorf("parm error")
	}

	deviceId, err := util.DeviceIdUid.Get()
	device.DeviceId = deviceId
	// 添加一个设备
	id, err := dao.DeviceDao.Add(device)
	if err != nil && id == -1 {
		return id, err
	}

	// 添加一个设备同步序列号记录
	if dao.DeviceAckDao.Add(device.DeviceId, 0) != nil {
		return -1, err
	}
	return id, nil
}

// 获取所有用户的所有的在线设备
func (*deviceService) ListOnlineByUserId(ctx context.Context, appId int64, userId int64) ([]model.Device, error) {
	// 从缓存中获取

	// 从数据库中获取
	devices, err := dao.DeviceDao.ListOnlineByUserId(appId, userId)
	if err != nil {
		return nil, err
	}
	// 存储到缓存中
	return devices, nil

}
