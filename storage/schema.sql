# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: activity-storage.crqripd5tnm8.eu-west-2.rds.amazonaws.com (MySQL 5.6.35-log)
# Database: activity_storage
# Generation Time: 2017-10-12 09:14:43 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE IF NOT EXISTS ws_aggregate_storage;

USE ws_aggregate_storage;

# Dump of table aggregation
# ------------------------------------------------------------

DROP TABLE IF EXISTS `aggregation`;

CREATE TABLE `aggregation` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `channel_id` int(11) DEFAULT NULL,
  `last_stream_session` timestamp NULL DEFAULT NULL,
  `last_channel_summary` timestamp NULL DEFAULT NULL,
  `last_subscription_summary` timestamp NULL DEFAULT NULL,
  `last_stream` timestamp NULL DEFAULT NULL,
  `last_users` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table channel
# ------------------------------------------------------------

DROP TABLE IF EXISTS `channel`;

CREATE TABLE `channel` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `mature` tinyint(4) DEFAULT NULL,
  `status` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `broadcaster_language` varchar(4) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `display_name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `game` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `language` varchar(4) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `_id` bigint(20) DEFAULT NULL,
  `name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `updated_at` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `partner` tinyint(4) DEFAULT NULL,
  `logo` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `video_banner` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `profile_banner` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `profile_banner_background_color` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `url` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `views` bigint(20) DEFAULT NULL,
  `followers` bigint(20) DEFAULT NULL,
  `broadcaster_type` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `stream_key` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `date_add` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table channel_video
# ------------------------------------------------------------

DROP TABLE IF EXISTS `channel_video`;

CREATE TABLE `channel_video` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `channel_id` bigint(11) DEFAULT NULL,
  `video_id` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `title` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description_html` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `broadcast_id` bigint(20) DEFAULT NULL,
  `broadcast_type` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `tag_list` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `views` int(11) DEFAULT NULL,
  `url` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `language` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `viewable` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `viewable_at` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `published_at` timestamp NULL DEFAULT NULL,
  `recorded_at` timestamp NULL DEFAULT NULL,
  `game` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `length` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `video_id` (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table credential
# ------------------------------------------------------------

DROP TABLE IF EXISTS `credential`;

CREATE TABLE `credential` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` varchar(64) DEFAULT NULL,
  `app_name` varchar(11) DEFAULT NULL,
  `channel_name` varchar(32) DEFAULT NULL,
  `channel_id` bigint(20) DEFAULT NULL,
  `access_token` varchar(256) NOT NULL DEFAULT '',
  `refresh_token` varchar(128) NOT NULL DEFAULT '',
  `scope` text DEFAULT NULL,
  `expires_in` int(11) DEFAULT NULL,
  `date_updated` timestamp NULL DEFAULT current_timestamp(),
  `email` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `bio` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `display_name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `logo` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `type` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `date_add` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
