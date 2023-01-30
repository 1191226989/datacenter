CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(100) NOT NULL DEFAULT '' COMMENT '登录密码',
  `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号码',
  `nickname` varchar(100) DEFAULT '' COMMENT '昵称',
  PRIMARY KEY (`id`),
  KEY `idx_username` (`username`) USING BTREE,
  KEY `idx_mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;