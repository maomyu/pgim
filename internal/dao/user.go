package dao

import (
	"database/sql"
	"fmt"

	"github.com/yuwe1/pgim/internal/model"
	"github.com/yuwe1/pgim/pkg/client/dbpool"
	"github.com/yuwe1/pgim/pkg/logger"
)

type userDao struct{}

var UserDao = new(userDao)

// 添加一个用户
func (*userDao) Add(user model.User) (int64, error) {
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
	// 使用INSERT INGORE语句，则会忽略导致错误的行，并将其余行插入到表中。
	result, err := db.Exec("insert ignore into user(app_id,user_id,nickname,sex,avatar_url,extra) values(?,?,?,?,?,?)",
		user.AppId, user.UserId, user.Nickname, user.Sex, user.AvatarUrl, user.Extra)
	if err != nil {
		return 0, fmt.Errorf("[%w]", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("[%w]", err)
	}
	return affected, nil
}

// 获取用户信息

func (*userDao) Get(appId int64, userId int64) (*model.User, error) {
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
	row := db.QueryRow("select nickname,sex,avatar_url,extra,create_time,update_time from user where app_id = ? and user_id = ?",
		appId, userId)
	user := model.User{
		AppId:  appId,
		UserId: userId,
	}
	err = row.Scan(&user.Nickname, &user.Sex, &user.AvatarUrl, &user.Extra, &user.CreateTime, &user.UpdateTime)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("%w", err)
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}
