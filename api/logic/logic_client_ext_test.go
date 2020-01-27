package logic

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/yuwe1/pgim/pb"
	"github.com/yuwe1/pgim/pkg/logger"
	"github.com/yuwe1/pgim/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func getLogicExtClient() pb.LogicClientExtClient {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pb.NewLogicClientExtClient(conn)
}
func getCtx() context.Context {
	token, _ := util.GetToken(1, 2, 3, time.Now().Add(1*time.Hour).Unix(), util.PublicKey)
	return metadata.NewOutgoingContext(context.TODO(), metadata.Pairs(
		"app_id", "1",
		"user_id", "2",
		"device_id", "3",
		"token", token,
		"request_id", strconv.FormatInt(time.Now().UnixNano(), 10)))
}
func TestLogicExtServer_RegisterDevice(t *testing.T) {
	resp, err := getLogicExtClient().RegisterDevice(context.TODO(), &pb.RegisterDeviceReq{
		AppId:         1,
		Type:          2,
		Brand:         "ios1",
		Model:         "ios 19991",
		SystemVersion: "1.0.10",
		SdkVersion:    "1.0.10",
	})
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Sugar.Info(resp)
}

func TestLogicExtServer_AddUser(t *testing.T) {
	resp, err := getLogicExtClient().AddUser(getCtx(),
		&pb.AddUserReq{
			User: &pb.User{
				Nicknamec: "10",
				Sex:       1,
				AvatarUrl: "10",
				Extra:     "10",
			},
		})
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Sugar.Info(resp)
}
