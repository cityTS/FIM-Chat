package model

import (
	"FIM-Chat/errors"
	"FIM-Chat/tools"
	"encoding/json"
	"fmt"
	"github.com/FengZhg/go_tools/logger"
	"github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	ID          int64           // 用户ID
	UnreadQueue *tools.Queue    // 未读消息队列(新连接推送全部队列消息，否则只推送新增加的消息)
	WIN         *websocket.Conn // Win端websocket长连接
	MAC         *websocket.Conn // MAC端websocket长连接
	Linux       *websocket.Conn // Linux端websocket长连接
}

var ClientCaching sync.Map

func DeleteWSConn(device int, id int64) {
	client, ok := ClientCaching.Load(id)
	if !ok {
		return
	}
	c := client.(*Client)
	switch device {
	case ClientLinux:
		closeConn(c.Linux)
		c.Linux = nil
	case ClientMAC:
		closeConn(c.MAC)
		c.MAC = nil
	case ClientWIN:
		closeConn(c.WIN)
		c.WIN = nil
	default:
		logger.Log.Error("无效的参数device: ", device)
	}
}
func AddWSConn(conn *websocket.Conn, device int, id int64) {
	client, ok := ClientCaching.Load(id)
	var c *Client
	if !ok {
		c = &Client{
			ID:          id,
			UnreadQueue: tools.NewQueue(),
			MAC:         nil,
			Linux:       nil,
			WIN:         nil,
		}
		ClientCaching.Store(id, c)
	} else {
		c = client.(*Client)
	}
	switch device {
	case ClientLinux:
		closeConn(c.Linux)
		c.Linux = conn
	case ClientWIN:
		closeConn(c.WIN)
		c.WIN = conn
	case ClientMAC:
		closeConn(c.MAC)
		c.MAC = conn
	default:
		logger.Log.Error("无效的参数")
	}
	fmt.Println(ClientCaching.Load(id))
}

// closeConn 关闭长连接
func closeConn(conn *websocket.Conn) {
	if conn == nil {
		return
	}
	m := Message{
		Type:   MessageTypeNotice,
		Notice: NoticeTypeLoginDuplicate,
	}
	mJson, err := json.Marshal(m)
	if err != nil {
		logger.Log.Warn(err)
	}
	err = conn.WriteMessage(websocket.TextMessage, mJson)
	if err != nil {
		logger.Log.Warn(err)
	}
	defer conn.Close()
}

// Broadcast 全平台广播信息
func Broadcast(id int64, m Message) error {
	client, ok := ClientCaching.Load(id)
	if !ok {
		return errors.NoLegalWSConnection
	}
	c, ok := client.(*Client)
	if !ok {
		return errors.ClientParseError
	}
	msg, err := json.Marshal(m)
	if err != nil {
		return err
	}
	if c.MAC != nil {
		err := c.MAC.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			logger.Log.Warn(err)
		}
	}
	if c.Linux != nil {
		err := c.Linux.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			logger.Log.Warn(err)
		}
	}
	if c.WIN != nil {
		err := c.WIN.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			logger.Log.Warn(err)
		}
	}
	return nil
}

// sendMessageMAC MAC平台推送
func sendMessageMAC(id int64, m Message) error {
	client, ok := ClientCaching.Load(id)
	if !ok {
		return errors.NoLegalWSConnection
	}
	c, ok := client.(Client)
	if !ok {
		return errors.ClientParseError
	}
	msg, err := json.Marshal(m)
	if err != nil {
		return err
	}
	if c.MAC != nil {
		err := c.MAC.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

// sendMessageWIN WIN平台推送
func sendMessageWIN(id int64, m Message) error {
	client, ok := ClientCaching.Load(id)
	if !ok {
		return errors.NoLegalWSConnection
	}
	c, ok := client.(Client)
	if !ok {
		return errors.ClientParseError
	}
	msg, err := json.Marshal(m)
	if err != nil {
		return err
	}
	if c.WIN != nil {
		err := c.WIN.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

// sendMessageLinux Linux平台推送
func sendMessageLinux(id int64, m Message) error {
	client, ok := ClientCaching.Load(id)
	if !ok {
		return errors.NoLegalWSConnection
	}
	c, ok := client.(Client)
	if !ok {
		return errors.ClientParseError
	}
	msg, err := json.Marshal(m)
	if err != nil {
		return err
	}
	if c.Linux != nil {
		err := c.Linux.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func AddUnreadMessage(id int64, message Message) {
	if client, ok := ClientCaching.Load(id); ok {
		c := client.(*Client)
		c.UnreadQueue.Add(message, message.ID)
	}
}
