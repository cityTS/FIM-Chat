package serve

import (
	"FIM-Chat/model"
	"FIM-Chat/tools"
	"encoding/json"
	"fmt"
	"github.com/FengZhg/go_tools/logger"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
)

// HandlerMessage 集中处理客户端的消息请求
func HandlerMessage(message model.Message, device int, id int64, conn *websocket.Conn) {
	switch message.Type {
	case model.MessageTypeNotice:
		handlerNotice(message, device, id, conn)
	case model.MessageTypeFile, model.MessageTypeImg, model.MessageTypeText:
		handlerMessage(message, device, id)
	}
}

// handlerNotice 处理通知类请求
func handlerNotice(message model.Message, device int, id int64, conn *websocket.Conn) {
	switch message.Notice {
	case model.NoticeTypeLogOut:
		logOut(device, id)
	case model.NoticePullUnread:
		pullingUnloadedMessages(message, conn, id)
	case model.NoticeRemoveFriend:
		removeFriend(message, id)

	}
}

func removeFriend(message model.Message, id int64) {
	model.DB.Model(&model.Message{}).Where("(from_user="+message.Content+"and to_user="+strconv.FormatInt(id, 10)+") or (from_user="+strconv.FormatInt(id, 10)+" to_user="+message.Content+")").Update("delete_time", time.Now().UnixMilli())
}

// pullingUnloadedMessages 拉取未加载的消息队列
func pullingUnloadedMessages(message model.Message, conn *websocket.Conn, id int64) {
	fmt.Println("开始拉取数据")
	var messageId, err = strconv.ParseInt(message.Content, 10, 64)
	if err != nil {
		fmt.Println(messageId, err)
		return
	}
	var messages []model.Message
	model.DB.Where("id > ? and (from_user = ? or to_user = ?) and delete_time = 0", messageId, id, id).Find(&messages)
	messageJson, err := json.Marshal(messages)
	fmt.Println("messageJson:", string(messageJson))
	if err != nil {
		logger.Log.Error(err)
		return
	}
	response := model.Message{
		Type:    model.MessageTypeNotice,
		Notice:  model.NoticePullUnread,
		Content: string(messageJson),
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	fmt.Println("responseJson:", string(responseJson))
	conn.WriteMessage(websocket.TextMessage, responseJson)
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
	message.Readed = true
	err = model.Broadcast(message.FromUser, message)
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
