package serve

import (
	"FIM-Chat/model"
	"encoding/json"
	"fmt"
	"github.com/FengZhg/go_tools/logger"
	"github.com/gorilla/websocket"
)

func AcceptMessage(conn *websocket.Conn, device int, id int64, ip string) {
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			model.DeleteWSConn(device, id)
			logger.Log.Info(err)
			break
		}
		if mt == websocket.TextMessage {
			var message model.Message
			json.Unmarshal(msg, &message)
			fmt.Println(message)
			message.IP = ip
			HandlerMessage(message, device, id)
		} else {
			logger.Log.Info("非法信息格式:", msg)
		}
	}
}
