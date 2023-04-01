CREATE TABLE IF NOT EXISTS `messages`
(
    `id`               BIGINT primary key comment '消息id',
    `type`             INTEGER      NOT NULL COMMENT 'text=1,img=2,file=3,notice=4',
    `notice`           INTEGER      NOT NULL comment '通知类型',
    `content`          VARCHAR(255) COMMENT '适用于text和notice类型消息',
    `img_url`          VARCHAR(255) COMMENT '使用户img类型',
    `file_type`        VARCHAR(255) COMMENT '适用于file类型',
    `file_name`        VARCHAR(255) COMMENT '适用于file类型',
    `file_url`         VARCHAR(255) COMMENT '适用于file类型',
    `from_user`        BIGINT       NOT NULL comment '用户id',
    `to_user`          BIGINT       NOT NULL comment '用户id',
    `to_user_type`     VARCHAR(255) NOT NULL COMMENT '用户/群',
    `is_group_message` BOOL         NOT NULL COMMENT '是否来自群消息\n\n如果User-A->Group-B，C in B,\n服务器将消息进行更改\nfromGroup = toUser\ntoUser=C\ntoUserType = user\nisGroupMessage = true',
    `from_group`       BIGINT       NOT NULL COMMENT '群号',
    `ip`               CHAR(15)     NOT NULL,
    `readed`           BOOL         NOT NULL COMMENT '是否已读',
    `create_time`      BIGINT       NOT NULL COMMENT '以服务器接收到消息为准',
    `delete_time`      BIGINT COMMENT '以服务器收到消息为准'
) ENGINE = InnoDB COMMENT '消息模型';