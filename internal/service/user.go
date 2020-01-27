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
