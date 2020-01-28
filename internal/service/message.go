package service

import (
	"context"

	"github.com/yuwe1/pgim/internal/dao"
	"github.com/yuwe1/pgim/internal/model"
	"github.com/yuwe1/pgim/pb"
	"github.com/yuwe1/pgim/pkg/grpclib"
	"github.com/yuwe1/pgim/pkg/util"
)

type messageService struct{}

var MessageService = new(messageService)

// Add 添加消息
func (*messageService) Add(ctx context.Context, message model.Message) error {
	return dao.MessageDao.Add("message", message)
}

func (*messageService) Send(ctx context.Context, sender model.Sender, in *pb.SendMessageReq) error {
	switch in.ReceiverType {
	case pb.ReceiverType_RT_USER:
		if sender.SenderType == pb.SenderType_ST_USER {
			err := MessageService.SendToFriend(ctx, sender, in)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 将消息发送到用户
func (*messageService) SendToFriend(ctx context.Context, sender model.Sender, in *pb.SendMessageReq) error {
	// 将消息发送至发送者

	// 将消息发送至接收者

	return nil
}

func (*messageService) SendToUser(ctx context.Context, sender model.Sender, toUserId, roomSeq int64, req pb.SendMessageReq) error {
	// 日志存储
	var (
		seq = roomSeq
		err error
	)
	if req.IsPersist {
		seq, err = SeqService.GetUserNext(ctx, sender.AppId, toUserId)
		if err != nil {
			return err
		}
		messageType, messageContent := model.PBToMesageBody(req.MessageBody)
		selfMessage := model.Message{
			// 发送者的appid
			AppId: sender.AppId,
			// 消息的所属类型
			ObjectType: model.MessageObjectTypeUser,
			// 接收者id,所属于类型的id，有可能时群组id也可能是用户id
			ObjectId:       toUserId,
			RequestId:      grpclib.GetCtxRequestId(ctx),
			SenderType:     int32(sender.SenderType),
			SenderId:       sender.SenderId,
			SenderDeviceId: sender.DeviceId,
			ReceiverType:   int32(req.ReceiverType),
			ReceiverId:     req.ReceiverId,
			ToUserIds:      model.FormatUserIds(req.ToUserId),
			// 消息的内容类型
			Type:     messageType,
			Content:  messageContent,
			Seq:      seq,
			SendTime: util.UnunixMilliTime(req.SendTime),
			// 消息的状态，未知，正常，撤回。。。
			Status: int32(pb.MessageStatus_MS_NORMAL),
		}
		// 将数据库中增加一条消息记录
		err = MessageService.Add(ctx, selfMessage)
		if err != nil {
			return err
		}
	}
	// 单条消息通知栏提醒
	messageItem := pb.MessageItem{
		RequestId:      grpclib.GetCtxRequestId(ctx),
		SenderType:     sender.SenderType,
		SenderId:       sender.SenderId,
		SenderDeviceId: sender.DeviceId,
		ReceiverType:   req.ReceiverType,
		ReceiverId:     req.ReceiverId,
		ToUserIds:      req.ToUserId,
		MessageBody:    req.MessageBody,
		Seq:            seq,
		SendTime:       req.SendTime,
		Status:         pb.MessageStatus_MS_NORMAL,
	}

	// 查询该用户的在线呢设备
	devices, err := DeviceService.ListOnlineByUserId(ctx, sender.AppId, toUserId)
	if err != nil {
		return err
	}
	for i := range devices {
		// 消息不需要投递到发送消息的设备（同步显示到其他的在线设备）
		if sender.DeviceId == devices[i].DeviceId {
			continue
		}
		// 将消息发送到设备
		err = MessageService.SendToDevice(ctx, devices[i], messageItem)
		if err != nil {
			return err
		}
	}
	return nil
}

// 将消息发送给设备
func (*messageService) SendToDevice(ctx context.Context, device model.Device, msgItem pb.MessageItem) error {
	if device.Status == model.DeviceOnLine {
		// message := pb.Message{Message: &msgItem}
		// _, err := rpc_cli.ConnectIntClient.DeliverMessage(grpclib.ContextWithAddr(ctx, device.ConnAddr), &pb.DeliverMessageReq{})
		// 调用连接层，发送消息

	}

	return nil
}
