CREATE TABLE `comment`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `creator_uid` int NOT NULL,
  `text` text NOT NULL,
  `video_id` int UNSIGNED NOT NULL,
  `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  `update_time` datetime NOT NULL,
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_video_id`(`video_id`) USING BTREE
);

CREATE TABLE `favorite`  (
  `id` int UNSIGNED NOT NULL,
  `creator_id` int UNSIGNED NOT NULL,
  `video_id` int UNSIGNED NOT NULL,
  `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  `update_time` datetime NOT NULL,
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_creator_id`(`creator_id`) USING BTREE,
  INDEX `idx_video_id`(`video_id`) USING BTREE
);

CREATE TABLE `follower`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `to_user_uid` int UNSIGNED NOT NULL COMMENT '对方id',
  `user_uid` int UNSIGNED NOT NULL,
  `create_time` datetime NOT NULL,
  `is_deleted` tinyint(1) UNSIGNED NOT NULL,
  `update_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_to_user_id`(`to_user_uid`) USING BTREE,
  INDEX `idx_user_id`(`user_uid`) USING BTREE
);

CREATE TABLE `message`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `reciver_id` int UNSIGNED NOT NULL,
  `listener_id` int UNSIGNED NOT NULL COMMENT '接收者',
  `text` text NOT NULL,
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `user`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `password` char(32) NOT NULL COMMENT 'MD5加盐处理',
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uk_name`(`name`) USING BTREE
);

CREATE TABLE `video`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `path` varchar(200) NOT NULL COMMENT '视频路径',
  `creator_id` int UNSIGNED NOT NULL,
  `create_time` datetime NOT NULL,
  `cover_path` varchar(200) NOT NULL COMMENT '封面路径',
  `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  `update_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_user_id`(`creator_id`) USING BTREE
);

