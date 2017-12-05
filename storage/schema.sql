# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.3.2-MariaDB-10.3.2+maria~jessie)
# Database: ws_aggregate_storage
# Generation Time: 2017-12-05 14:06:20 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table api_channels
# ------------------------------------------------------------

DROP TABLE IF EXISTS `api_channels`;

CREATE TABLE `api_channels` (
  `meta_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `meta_date_add` timestamp NULL DEFAULT NULL,
  `id` bigint(20) DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `mature` tinyint(1) DEFAULT NULL,
  `status` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `broadcaster_language` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `display_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `game` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `language` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `partner` tinyint(1) DEFAULT NULL,
  `logo` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `video_banner` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `profile_banner` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `profile_banner_bg_color` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `views` bigint(20) DEFAULT NULL,
  `followers` bigint(20) DEFAULT NULL,
  `broadcaster_type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `stream_key` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`meta_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table api_credentials
# ------------------------------------------------------------

DROP TABLE IF EXISTS `api_credentials`;

CREATE TABLE `api_credentials` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `meta_date_add` timestamp NULL DEFAULT NULL,
  `meta_date_update` timestamp NULL DEFAULT NULL,
  `channel_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `channel_id` bigint(20) DEFAULT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `access_token` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `refresh_token` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `expires_in` bigint(20) DEFAULT NULL,
  `scopes` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_api_credentials_channel_id` (`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table api_users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `api_users`;

CREATE TABLE `api_users` (
  `meta_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `meta_date_add` timestamp NULL DEFAULT NULL,
  `meta_date_update` timestamp NULL DEFAULT NULL,
  `id` bigint(20) DEFAULT NULL,
  `display_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `bio` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `logo` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`meta_id`),
  UNIQUE KEY `uix_api_users_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table api_videos
# ------------------------------------------------------------

DROP TABLE IF EXISTS `api_videos`;

CREATE TABLE `api_videos` (
  `meta_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `meta_date_add` timestamp NULL DEFAULT NULL,
  `id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description_html` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `broadcast_id` bigint(20) DEFAULT NULL,
  `broadcast_type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `tag_list` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `views` bigint(20) DEFAULT NULL,
  `url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `language` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `viewable` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `viewable_at` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `published_at` timestamp NULL DEFAULT NULL,
  `recorded_at` timestamp NULL DEFAULT NULL,
  `game` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `length` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`meta_id`),
  UNIQUE KEY `uix_api_videos_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table precomputed_channels
# ------------------------------------------------------------

DROP TABLE IF EXISTS `precomputed_channels`;

CREATE TABLE `precomputed_channels` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `meta_date_add` timestamp NULL DEFAULT NULL,
  `channel_id` bigint(20) DEFAULT NULL,
  `channel_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `avg_ccv` bigint(20) DEFAULT NULL,
  `max_ccv` bigint(20) DEFAULT NULL,
  `air_time` bigint(20) DEFAULT NULL,
  `seconds_watched` bigint(20) DEFAULT NULL,
  `primary_game` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `partner` tinyint(1) DEFAULT NULL,
  `mature` tinyint(1) DEFAULT NULL,
  `language` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `views` bigint(20) DEFAULT NULL,
  `followers` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
