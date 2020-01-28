package model

import (
	"encoding/json"

	"github.com/yuwe1/pgim/pb"
	"github.com/yuwe1/pgim/pkg/logger"
	"github.com/yuwe1/pgim/pkg/util"
)

// 将messageBody转换成类型和内容
func PBToMesageBody(pbBody *pb.MessageBody) (int, string) {
	if pbBody.MessageType == pb.MessageType_MT_UNKNOWN {
		// 未知消息
		logger.Logger.Error("error message type")
		return 0, ""
	}
	// 定义一个内容
	var content interface{}
	// 根据messageType寻找到应该转换的内容
	switch pbBody.MessageType {
	case pb.MessageType_MT_TEXT:
		// 文本消息
		content = pbBody.MessageContent.GetText()
	}
	// 开始进行转换成字节类型
	bytes, err := json.Marshal(content)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, ""
	}
	return int(pbBody.MessageType), util.Bytes2str(bytes)
}
