package server

import (
	"FIM-Chat/server/serve"
	"github.com/FengZhg/go_tools/logger"
	"github.com/gorilla/websocket"
)

type HubConn struct {
	Conn   *websocket.Conn
	Device int
	Id     int64
	IP     string
}

var (
	HubConnCaching chan HubConn
)

const (
	// HubCachingSize 默认缓存大小
	hubCachingSize = 1000
)

func init() {
	HubConnCaching = make(chan HubConn, hubCachingSize)
	go ListenServer()
}

func ListenServer() {
	logger.Log.Info("数据中心监听服务已启动")
	for {
		select {
		case c := <-HubConnCaching:
			go serve.AcceptMessage(c.Conn, c.Device, c.Id, c.IP)
		}
	}
	defer logger.Log.Info("数据中心监听服务已关闭")
}
