package dao

import (
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
