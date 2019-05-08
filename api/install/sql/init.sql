# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.25)
# Database: gold
# Generation Time: 2019-05-06 06:33:18 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table function_services
# ------------------------------------------------------------

DROP TABLE IF EXISTS `function_services`;

CREATE TABLE `function_services` (
                                   `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
                                   `creator_id` bigint(11) NOT NULL,
                                   `creator_name` varchar(100) NOT NULL DEFAULT '',
                                   `service_name` varchar(100) NOT NULL DEFAULT '',
                                   `git_repo` varchar(255) NOT NULL DEFAULT '',
                                   `git_branch` varchar(255) NOT NULL DEFAULT '',
                                   `git_head` varchar(100) DEFAULT '',
                                   `git_maintainer` varchar(100) NOT NULL DEFAULT '',
                                   `status` varchar(20) NOT NULL DEFAULT 'created',
                                   `last_operation` bigint(11) NOT NULL DEFAULT '0',
                                   `add_on` varchar(255) DEFAULT NULL,
                                   `min_instance` int(11) NOT NULL DEFAULT '1',
                                   `max_instance` int(11) NOT NULL DEFAULT '1',
                                   `created_at` bigint(13) DEFAULT NULL,
                                   `updated_at` bigint(13) DEFAULT NULL,
                                   PRIMARY KEY (`id`),
                                   KEY `service_name` (`service_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table operate_logs
# ------------------------------------------------------------

DROP TABLE IF EXISTS `operate_logs`;

CREATE TABLE `operate_logs` (
                              `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
                              `service_id` bigint(11) NOT NULL,
                              `operator_id` bigint(11) NOT NULL COMMENT '操作执行者id',
                              `type` varchar(20) NOT NULL DEFAULT '' COMMENT '操作类型',
                              `start` bigint(13) NOT NULL COMMENT '操作开始时间',
                              `update` bigint(13) DEFAULT NULL COMMENT '操作更新时间',
                              `end` bigint(13) DEFAULT NULL COMMENT '操作结束时间',
                              `current_action` varchar(100) DEFAULT NULL COMMENT '当前执行动作',
                              `log` text COMMENT '相关输出日志',
                              `origin_branch` varchar(100) DEFAULT NULL COMMENT '变更前分支',
                              `origin_version` varchar(100) DEFAULT NULL COMMENT '变更前版本号',
                              `target_branch` varchar(100) DEFAULT NULL COMMENT '变更后分支',
                              `target_version` varchar(100) DEFAULT NULL COMMENT '变更后版本号',
                              `add_on` varchar(255) DEFAULT NULL COMMENT '冗余',
                              `created_at` bigint(13) DEFAULT NULL,
                              `updated_at` bigint(13) DEFAULT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
                       `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
                       `name` varchar(100) NOT NULL DEFAULT '',
                       `email` varchar(100) NOT NULL DEFAULT '',
                       `password` varchar(100) NOT NULL,
                       `created_at` bigint(13) DEFAULT NULL,
                       `updated_at` bigint(13) DEFAULT NULL,
                       `add_on` varchar(255) DEFAULT NULL,
                       PRIMARY KEY (`id`),
                       UNIQUE KEY `name` (`name`),
                       UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
