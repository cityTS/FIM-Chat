package model

// 客户端标识

var (
	ClientWIN   = 1
	ClientMAC   = 2
	ClientLinux = 3
)

// 消息标识

var (
	MessageTypeText   = 1
	MessageTypeImg    = 2
	MessageTypeFile   = 3
	MessageTypeNotice = 4
)

// 通知类型

var (
	NoticeTypeLoginDuplicate       = 1 // 重复登陆
	NoticeTypeInvalidRequestFormat = 2 // 无效的请求格式
	NoticeTypeLogOut               = 3 // 退出登录
	NoticePullUnread               = 4 // 拉取未读队列
)
