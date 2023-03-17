package model

import "FIM-Chat/tools"

type Client struct {
	ID          int64       // 用户ID
	UnreadQueue tools.Queue // 未读消息队列
}
