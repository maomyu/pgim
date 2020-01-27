package dao

import (
	"fmt"
	"github.com/yuwe1/pgim/internal/model"
	"github.com/yuwe1/pgim/pkg/client/dbpool"
	"github.com/yuwe1/pgim/pkg/logger"
)

type deviceDao struct{}

var DeviceDao = new(deviceDao)

func (*deviceDao) Add(device model.Device) (int64, error) {
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
	sql, err := db.Exec(`insert into device(device_id,app_id,type,brand,model,system_version,sdk_version,status,conn_addr) 
	values(?,?,?,?,?,?,?,?,?)`,
		device.DeviceId, device.AppId, device.Type, device.Brand, device.Model, device.SystemVersion, device.SDKVersion, device.Status, "")
	if err != nil {
		return -1, fmt.Errorf("[%w]", err)
	}
	id, _ := sql.LastInsertId()
	return id, nil
}
