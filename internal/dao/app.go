package dao

import (
	"database/sql"
	"fmt"

	"github.com/yuwe1/pgim/internal/model"
	"github.com/yuwe1/pgim/pkg/client/dbpool"
	"github.com/yuwe1/pgim/pkg/logger"
)

type appDao struct {
}

var AppDao = new(appDao)

func (*appDao) Get(appId int64) (*model.App, error) {
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
	var app model.App

	err = db.QueryRow("select id,name,private_key,create_time,update_time from app where id = ?", appId).Scan(
		&app.Id, &app.Name, &app.PrivateKey, &app.CreateTime, &app.UpdateTime)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("database error [%w]", err)
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &app, nil
}
