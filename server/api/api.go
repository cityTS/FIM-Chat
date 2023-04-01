package api

import (
	"FIM-Chat/model"
	"FIM-Chat/server"
	"github.com/FengZhg/go_tools/logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// EstablishWSConnection 建立Websocket连接
func EstablishWSConnection(c *gin.Context) {
	device := c.Query("device")
	if device == "" {
		logger.Log.Info("缺少客户端信息")
		c.JSON(http.StatusOK, model.Message{
			Type:   model.MessageTypeNotice,
			Notice: model.NoticeTypeInvalidRequestFormat,
		})
		return
	}
	id := c.Query("userId")
	if id == "" {
		logger.Log.Info("缺少客户端信息")
		c.JSON(http.StatusOK, model.Message{
			Type:   model.MessageTypeNotice,
			Notice: model.NoticeTypeInvalidRequestFormat,
		})
		return
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Log.Error("无法升级HTTP请求: ", err.Error())
		return
	}
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logger.Log.Info("缺少客户端信息")
		c.JSON(http.StatusOK, model.Message{
			Type:   model.MessageTypeNotice,
			Notice: model.NoticeTypeInvalidRequestFormat,
		})
		defer ws.Close()
		return
	}
	d, err := strconv.Atoi(device)
	if err != nil {
		logger.Log.Info("缺少客户端信息")
		c.JSON(http.StatusOK, model.Message{
			Type:   model.MessageTypeNotice,
			Notice: model.NoticeTypeInvalidRequestFormat,
		})
		defer ws.Close()
		return
	}
	// 注册服务
	model.AddWSConn(ws, d, id64)
	server.HubConnCaching <- server.HubConn{
		Conn:   ws,
		Device: d,
		Id:     id64,
		IP:     getRealIp(c),
	}
	logger.Log.Info("用户:", id, " 在", getDevice(d), "设备上登录")
}
func getRealIp(c *gin.Context) string {
	if c.Request.Header.Get("X-Forwarded-For") != "" {
		return c.Request.Header.Get("X-Forwarded-For")
	} else {
		return c.RemoteIP()
	}
}
func getDevice(device int) string {
	switch device {
	case model.ClientMAC:
		return "MAC端"
	case model.ClientWIN:
		return "WIN端"
	case model.ClientLinux:
		return "Linux端"
	default:
		return "未知设备"
	}
}
