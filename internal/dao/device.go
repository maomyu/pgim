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

// 获取所有的在线设备
func (*deviceDao) ListOnlineByUserId(appId, userId int64) ([]model.Device, error) {
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
	rows, err := db.Query(
		`select device_id,type,brand,model,system_version,sdk_version,status,conn_addr,create_time,update_time from device where app_id = ? and user_id = ? and status = ?`,
		appId, userId, model.DeviceOnLine)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	devices := make([]model.Device, 0, 5)
	for rows.Next() {
		device := new(model.Device)
		err = rows.Scan(&device.DeviceId, &device.Type, &device.Brand, &device.Model, &device.SystemVersion, &device.SDKVersion,
			&device.Status, &device.ConnAddr, &device.CreateTime, &device.UpdateTime)
		if err != nil {
			logger.Sugar.Error(err)
			return nil, err
		}
		devices = append(devices, *device)
	}
	return devices, nil
}
