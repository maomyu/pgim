package dao

import (
	"fmt"
	"github.com/yuwe1/pgim/pkg/client/dbpool"
	"github.com/yuwe1/pgim/pkg/logger"
)

type deviceAckDao struct{}

var DeviceAckDao = new(deviceAckDao)

// 添加设备同步序列号
func (*deviceAckDao) Add(deviceId int64, ack int64) error {
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
	_, err = db.Exec("insert into device_ack(device_id,ack) values(?,?)", deviceId, ack)
	if err != nil {
		return fmt.Errorf("[%w]", err)
	}
	return nil
}
