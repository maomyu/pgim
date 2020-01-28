package logic

import (
	"context"
	"fmt"

	"github.com/yuwe1/pgim/internal/model"
	"github.com/yuwe1/pgim/internal/service"
	"github.com/yuwe1/pgim/pb"
	"github.com/yuwe1/pgim/pkg/gerrors"
	"github.com/yuwe1/pgim/pkg/grpclib"
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

	if device.AppId == 0 || device.Type == 0 || device.Brand == "" || device.Model == "" || device.SystemVersion == "" || device.SDKVersion == "" {
		return nil, fmt.Errorf("bad param")
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

func (*LogicClientExt) AddUser(ctx context.Context, req *pb.AddUserReq) (*pb.AddUserResp, error) {
	// 获取用户信息
	appId, userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return &pb.AddUserResp{}, err
	}
	user := model.User{
		AppId:     appId,
		UserId:    userId,
		Nickname:  req.User.Nicknamec,
		Sex:       req.User.Sex,
		AvatarUrl: req.User.AvatarUrl,
		Extra:     req.User.Extra,
	}

	return &pb.AddUserResp{}, service.UserService.Add(ctx, user)
}
func (*LogicClientExt) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserResp, error) {
	appId, _, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return &pb.GetUserResp{}, err
	}
	user, err := service.UserService.Get(ctx, appId, req.UserId)
	if err != nil {
		return &pb.GetUserResp{}, nil
	}
	if user == nil {
		return nil, fmt.Errorf("[%w]", gerrors.ErrUserNotExist)
	}

	pbUser := pb.User{
		UserId:     user.UserId,
		Nicknamec:  user.Nickname,
		Sex:        user.Sex,
		AvatarUrl:  user.AvatarUrl,
		Extra:      user.Extra,
		CreateTime: user.CreateTime.Unix(),
		UpdateTime: user.UpdateTime.Unix(),
	}
	return &pb.GetUserResp{User: &pbUser}, nil
}

//发送消息
func (*LogicClientExt) SendMessage(ctx context.Context, req *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	appId, userId, deviceId, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}
	sender := model.Sender{
		AppId:      appId,
		SenderType: pb.SenderType_ST_USER,
		SenderId:   userId,
		DeviceId:   deviceId,
	}
	err = service.MessageService.Send(ctx, sender, req)
}
