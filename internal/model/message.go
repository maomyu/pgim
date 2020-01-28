package model

import "time"

import "strings"

import "strconv"

const (
	MessageObjectTypeUser  = 1 // 用户
	MessageObjectTypeGroup = 2 // 群组
)

// Message 消息
type Message struct {
	Id             int64     // 自增主键
	AppId          int64     // appId
	ObjectType     int       // 所属类型
	ObjectId       int64     // 所属类型id
	RequestId      int64     // 请求id
	SenderType     int32     // 发送者类型
	SenderId       int64     // 发送者账户id
	SenderDeviceId int64     // 发送者设备id
	ReceiverType   int32     // 接收者账户id
	ReceiverId     int64     // 接收者id,如果是单聊信息，则为user_id，如果是群组消息，则为group_id
	ToUserIds      string    // 需要@的用户id列表，多个用户用，隔开
	Type           int       // 消息类型
	Content        string    // 消息内容
	Seq            int64     // 消息同步序列
	SendTime       time.Time // 消息发送时间
	Status         int32     // 创建时间
}

// 对@的用户Id格式进行转换
func FormatUserIds(userId []int64) string {
	build := strings.Builder{}
	for i, v := range userId {
		build.WriteString(strconv.FormatInt(v, 10))
		if i != len(userId)-1 {
			build.WriteString(",")
		}
	}
	return build.String()
}
