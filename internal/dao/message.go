package dao

import (
	"fmt"

	"github.com/yuwe1/pgim/internal/model"
	"github.com/yuwe1/pgim/pkg/client/dbpool"
	"github.com/yuwe1/pgim/pkg/logger"
)

type messageDao struct {
}

var MessageDao = new(messageDao)

// 插入一条消息
func (*messageDao) Add(tableName string, message model.Message) error {
	session, err, p, c := dbpool.GetSession()
	defer func() {
		if session != nil {
			session.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	db := session.DB
	sql := fmt.Sprintf(`insert into %s(app_id,object_type,object_id,request_id,sender_type,sender_id,sender_device_id,receiver_type,receiver_id,to_user_ids,type,content,seq,send_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, tableName)
	_, err = db.Exec(sql, message.AppId, message.ObjectType, message.ObjectId, message.RequestId, message.SenderType, message.SenderId, message.SenderDeviceId,
		message.ReceiverType, message.ReceiverId, message.ToUserIds, message.Type, message.Content, message.Seq, message.SendTime)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
