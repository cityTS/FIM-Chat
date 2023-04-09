package model

type Message struct {
	ID             int64  `json:"id" gorm:"column:id"`                           // 消息唯一ID
	Content        string `json:"content" gorm:"column:content"`                 // 消息内容，适用于text和notice类型消息
	CreateTime     int64  `json:"createTime" gorm:"column:create_time"`          // 发送时间，以服务器接收到消息为准
	DeleteTime     int64  `json:"deleteTime" gorm:"column:delete_time"`          // 撤销消息时间或失效时间（如删除好友、退出群聊，则失去此消息的浏览权限），以服务器收到消息为准
	FileName       string `json:"fileName" gorm:"column:file_name"`              // 文件名称，适用于file类型
	FileType       string `json:"fileType" gorm:"column:file_type"`              // 文件类型(后缀)，适用于file类型
	FileURL        string `json:"fileUrl" gorm:"column:file_url"`                // 文件链接，适用于file类型
	FromGroup      int64  `json:"fromGroup" gorm:"column:from_group"`            // 群号
	FromUser       int64  `json:"fromUser" gorm:"column:from_user"`              // 发送者id
	ImgURL         string `json:"imgUrl" gorm:"column:img_url"`                  // 图片链接，使用户img类型
	IP             string `json:"ip" gorm:"column:ip"`                           // 发送者ip
	IsGroupMessage bool   `json:"isGroupMessage" gorm:"column:is_group_message"` // 是否来自群消息; ; 如果User-A->Group-B，C in B,; 服务器将消息进行更改; fromGroup = toUser; toUser=C; toUserType = user; isGroupMessage = true
	Readed         bool   `json:"readed" gorm:"column:readed"`                   // 是否已读
	ToUser         int64  `json:"toUser" gorm:"column:to_user"`                  // 接收者id
	Type           int    `json:"type" gorm:"column:type"`                       // 消息类型，text=1,img=2,file=3,notice=4
	Notice         int    `json:"notice" gorm:"column:notice"`                   // 通知类型 1：异地同端登录
}
