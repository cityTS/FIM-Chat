package ws

import (
	"FIM-Chat/server/api"
	"FIM-Chat/server/mid"
	"github.com/gin-gonic/gin"
)

func init() {
	logon()
}

func logon() {
	r := gin.Default()
	r.Use(mid.Cors())

	r.GET("/api/addWSConn", api.EstablishWSConnection)

	r.Run(":4399")
}
