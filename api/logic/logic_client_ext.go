package logic

import (
	"context"

	"github.com/yuwe1/pgim/internal/model"
	"github.com/yuwe1/pgim/internal/service"
	"github.com/yuwe1/pgim/pb"
)

type LogicClientExt struct {
}

func (*LogicClientExt) RegisterDevice(ctx context.Context, req *pb.RegisterDeviceReq) (*pb.RegisterDeviceResp, error) {
	device := model.Device{
		AppId:         req.AppId,
		Type:          req.Type,
		Brand:         req.Brand,
		Model:         req.Model,
		SystemVersion: req.SystemVersion,
		SDKVersion:    req.SdkVersion,
	}

	// 判断是否出现参数错误
	if device.AppId == 0 ||
		device.Type == 0 || device.Brand == "" ||
		device.Model == "" || device.SystemVersion == "" ||
		device.SDKVersion == "" {
		// 定义一个错误
		return nil, nil
	}
	// 调用service注册逻辑
	id, err := service.DeviceService.Register(ctx, device)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterDeviceResp{DeviceId: id}, nil
}
