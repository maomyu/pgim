package service

import (
	"context"
	"fmt"

	"github.com/yuwe1/pgim/internal/dao"
	"github.com/yuwe1/pgim/internal/model"
)

type userService struct{}

var UserService = new(userService)

// 添加用户,添加用户的消息序列号
func (*userService) Add(ctx context.Context, user model.User) error {
	affected, err := dao.UserDao.Add(user)
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("用户已经存在")
	}
	return nil
}

// 获取用户信息
func (*userService) Get(ctx context.Context, appId int64, userId int64) (*model.User, error) {
	// 根据appId和userId获取用户信息的缓存

	// 从数据库中获取
	user, err := dao.UserDao.Get(appId, userId)
	if err != nil {
		return nil, err
	}
	if user != nil {
		// 将用户信息存储到缓存中

	}
	return user, err
}
