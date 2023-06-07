-- sys
CREATE TABLE `sys_user` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
    `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
    `nick_name` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称',
    `avatar` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
    `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
    `salt` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '加密盐',
    `email` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '邮箱',
    `mobile` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机号',
    `status` tinyint NOT NULL COMMENT '状态  0：禁用   1：正常',
    `dept_id` bigint NOT NULL COMMENT '机构ID',
    `create_by` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '创建人',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_by` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '更新人',
    `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `del_flag` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除  -1：已删除  0：正常',
    `job_id` int NOT NULL COMMENT '岗位Id',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户管理'

CREATE TABLE `sys_login_log` (
 `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
 `user_name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
 `status` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）',
 `ip` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'IP地址',
 `create_by` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '创建人',
 `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统登录日志'

