package serve

import (
	"FIM-Chat/model"
	"encoding/json"
	"github.com/FengZhg/go_tools/logger"
	"github.com/gorilla/websocket"
)

func AcceptMessage(conn *websocket.Conn, device int, id int64, ip string) {
	logger.Log.Info("开始监听用户:", id)
	num := 1
	for {
		mt, msg, err := conn.ReadMessage()
		logger.Log.Info("用户:", id, " 数据:", num)
		num++
		if err != nil {
			model.DeleteWSConn(device, id)
			logger.Log.Info(err)
			logger.Log.Info("结束监听用户:", id)
			break
		}
		if mt == websocket.TextMessage {
			var message model.Message
			err := json.Unmarshal(msg, &message)
			if err != nil {
				logger.Log.Error(err)
			}
			message.IP = ip
			HandlerMessage(message, device, id, conn)
		} else {
			logger.Log.Info("非法信息格式:", msg)
		}
	}
}
