package serve

import (
	"FIM-Chat/model"
	"FIM-Chat/tools"
	"github.com/FengZhg/go_tools/logger"
	"time"
)

// HandlerMessage 集中处理客户端的消息请求
func HandlerMessage(message model.Message, device int, id int64) {
	switch message.Type {
	case model.MessageTypeNotice:
		handlerNotice(message, device, id)
	case model.MessageTypeFile, model.MessageTypeImg, model.MessageTypeText:
		handlerMessage(message, device, id)
	}
}

// handlerNotice 处理通知类请求
func handlerNotice(message model.Message, device int, id int64) {
	switch message.Notice {
	case model.NoticeTypeLogOut:
		logOut(device, id)
	case model.NoticePullUnread:
		//TODO 新登录拉取未读消息队列
	}
}

// logOut 退出登录
func logOut(device int, id int64) {
	model.DeleteWSConn(device, id)
}

// handlerMessage 处理消息记录
func handlerMessage(message model.Message, device int, id int64) {
	switch message.IsGroupMessage {
	case true:
		handlerGroupMessage(message)
	default:
		handlerSingleMessage(message)
	}
}

// handlerSingleMessage 处理单对单消息
func handlerSingleMessage(message model.Message) {
	// 利用SnowFlakes算法，给消息编号
	message.ID = tools.Snow.NextVal()
	message.CreateTime = time.Now().UnixMilli()
	if !message.Persistence() {
		return
	}

	err := model.Broadcast(message.ToUser, message)
	if err != nil {
		logger.Log.Info(err)
	}
}

// TODO 群管理
func handlerGroupMessage(message model.Message) {
	createGroupMessage(message)
}

func createGroupMessage(message model.Message) []model.Message {
	//TODO 拉取群组成员：考虑使用缓存 缓存更新策略：旁路缓存
	return nil
}
