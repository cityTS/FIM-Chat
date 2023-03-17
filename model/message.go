package model

type Message struct {
	ID             int64  `json:"id"`             // 消息唯一ID
	Content        string `json:"content"`        // 消息内容，适用于text和notice类型消息
	CreateTime     int64  `json:"createTime"`     // 发送时间，以服务器接收到消息为准
	DeleteTime     int64  `json:"deleteTime"`     // 撤销消息时间或失效时间（如删除好友、退出群聊，则失去此消息的浏览权限），以服务器收到消息为准
	FileName       string `json:"fileName"`       // 文件名称，适用于file类型
	FileType       string `json:"fileType"`       // 文件类型(后缀)，适用于file类型
	FileURL        string `json:"fileUrl"`        // 文件链接，适用于file类型
	FromGroup      string `json:"fromGroup"`      // 群号
	FromUser       string `json:"fromUser"`       // 发送者id
	ImgURL         string `json:"imgUrl"`         // 图片链接，使用户img类型
	IP             string `json:"ip"`             // 发送者ip
	IsGroupMessage bool   `json:"isGroupMessage"` // 是否来自群消息; ; 如果User-A->Group-B，C in B,; 服务器将消息进行更改; fromGroup = toUser; toUser=C; toUserType = user; isGroupMessage = true
	Readed         bool   `json:"readed"`         // 是否已读
	ToUser         string `json:"toUser"`         // 接收者id
	ToUserType     string `json:"toUserType"`     // 接收者类型，用户/群
	Type           string `json:"type"`           // 消息类型，text,img,file,notice
}
