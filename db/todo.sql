# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.17)
# Database: todo
# Generation Time: 2017-08-17 08:02:07 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table tasks
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tasks`;

CREATE TABLE `tasks` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `content` varchar(255) DEFAULT NULL COMMENT 'task内容',
  `status` int(11) DEFAULT NULL COMMENT '是否完成',
  `degree` int(11) DEFAULT NULL COMMENT '重要程度',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tasks` WRITE;
/*!40000 ALTER TABLE `tasks` DISABLE KEYS */;

INSERT INTO `tasks` (`id`, `content`, `status`, `degree`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'hi',1,1,'2017-08-17 08:27:49','2017-08-17 08:27:52',NULL),
	(2,'11',1,1,'2017-08-17 08:29:34','2017-08-17 10:01:28',NULL),
	(3,'22',0,1,'2017-08-17 08:29:41','2017-08-17 08:31:09',NULL),
	(4,'33',0,1,'2017-08-17 08:29:42','2017-08-17 08:29:42',NULL),
	(5,'44',0,1,'2017-08-17 08:29:45','2017-08-17 08:29:45',NULL),
	(6,'44',0,1,'2017-08-17 08:29:47','2017-08-17 10:01:10',NULL),
	(7,'555',1,1,'2017-08-17 08:29:49','2017-08-17 14:20:04',NULL),
	(8,'55',0,1,'2017-08-17 08:30:00','2017-08-17 10:01:05',NULL),
	(9,'9',0,1,'2017-08-17 08:30:38','2017-08-17 08:30:38',NULL),
	(10,'10',0,1,'2017-08-17 08:30:40','2017-08-17 08:30:40',NULL),
	(11,'11',1,1,'2017-08-17 08:30:46','2017-08-17 10:01:22',NULL);

/*!40000 ALTER TABLE `tasks` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `last_login_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `email`, `password`, `created_at`, `updated_at`, `deleted_at`, `last_login_time`)
VALUES
	(4,'haohongfan@123.com','123456','2017-07-30 19:20:10','2017-08-01 10:15:19',NULL,'2017-08-01 10:15:19'),
	(5,'haohongfan@126.com','123456','2017-07-30 21:46:45','2017-08-17 15:55:43',NULL,'2017-08-17 15:55:43'),
	(6,'haohongfan@ling.ai','111111','2017-08-17 15:56:07','2017-08-17 15:56:07',NULL,'2017-08-17 15:56:07');

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
