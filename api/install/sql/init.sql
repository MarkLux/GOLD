CREATE DATABASE IF NOT EXISTS `gold`;

CREATE TABLE `function_services` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `creator_id` bigint(11) NOT NULL,
  `creator_name` varchar(100) NOT NULL DEFAULT '',
  `service_name` varchar(100) NOT NULL DEFAULT '',
  `git_remote` varchar(255) NOT NULL DEFAULT '',
  `git_branch` varchar(255) NOT NULL DEFAULT '',
  `git_head` varchar(100) DEFAULT '',
  `status` varchar(20) NOT NULL DEFAULT 'created',
  `last_operation` bigint(11) NOT NULL DEFAULT '0',
  `add_on` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `service_name` (`service_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `operate_logs` (
                              `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
                              `service_id` bigint(11) NOT NULL,
                              `operator_id` bigint(11) NOT NULL COMMENT '操作执行者id',
                              `type` varchar(20) NOT NULL DEFAULT '',
                              `start` bigint(13) NOT NULL COMMENT '操作开始时间',
                              `update` bigint(13) DEFAULT NULL COMMENT '操作更新时间',
                              `end` bigint(13) DEFAULT NULL COMMENT '操作结束时间',
                              `current_action` varchar(100) DEFAULT NULL COMMENT '当前执行动作',
                              `log` text COMMENT '相关输出日志',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `users` (
                       `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
                       `name` varchar(100) NOT NULL DEFAULT '',
                       `email` varchar(100) NOT NULL DEFAULT '',
                       `created_at` timestamp NULL DEFAULT NULL,
                       `updated_at` timestamp NULL DEFAULT NULL,
                       `add_on` varchar(255) DEFAULT NULL,
                       PRIMARY KEY (`id`),
                       UNIQUE KEY `name` (`name`),
                       UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;