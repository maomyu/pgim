package logic

import (
	"context"
	"fmt"
	"testing"

	"github.com/yuwe1/pgim/pb"
	"github.com/yuwe1/pgim/pkg/logger"
	"google.golang.org/grpc"
)

func getLogicExtClient() pb.LogicClientExtClient {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pb.NewLogicClientExtClient(conn)
}

func TestLogicExtServer_RegisterDevice(t *testing.T) {
	resp, err := getLogicExtClient().RegisterDevice(context.TODO(), &pb.RegisterDeviceReq{
		AppId:         1,
		Type:          1,
		Brand:         "ios",
		Model:         "ios 1999",
		SystemVersion: "1.0.0",
		SdkVersion:    "1.0.0",
	})
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Sugar.Info(resp)
}
