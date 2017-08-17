# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.17)
# Database: todo
# Generation Time: 2017-08-17 11:36:02 +0000
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
  `user_id` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users` (`user_id`),
  CONSTRAINT `fk_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tasks` WRITE;
/*!40000 ALTER TABLE `tasks` DISABLE KEYS */;

INSERT INTO `tasks` (`id`, `content`, `status`, `degree`, `created_at`, `updated_at`, `deleted_at`, `user_id`)
VALUES
	(2,'11',0,1,'2017-08-17 17:37:31','2017-08-17 17:37:31',NULL,1),
	(3,'111',0,1,'2017-08-17 17:50:47','2017-08-17 17:50:47',NULL,1),
	(4,'33',0,1,'2017-08-17 17:50:53','2017-08-17 17:50:53',NULL,1),
	(5,'11',0,1,'2017-08-17 17:50:57','2017-08-17 17:50:57',NULL,1),
	(6,'33',0,1,'2017-08-17 17:50:58','2017-08-17 17:50:58',NULL,1),
	(7,'44',0,1,'2017-08-17 17:51:00','2017-08-17 17:51:00',NULL,1),
	(8,'555',0,1,'2017-08-17 17:51:02','2017-08-17 17:51:02',NULL,1),
	(9,'222',0,1,'2017-08-17 17:51:04','2017-08-17 17:51:04',NULL,1),
	(10,'1121',0,1,'2017-08-17 17:51:07','2017-08-17 17:51:07',NULL,1),
	(11,'332322',0,1,'2017-08-17 17:51:10','2017-08-17 17:51:10',NULL,1),
	(12,'111',0,1,'2017-08-17 17:51:12','2017-08-17 17:51:12',NULL,1),
	(13,'23121',0,1,'2017-08-17 18:14:39','2017-08-17 18:14:39',NULL,1);

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
	(1,'haohongfan@ling.ai','123456','2017-08-17 17:32:50','2017-08-17 19:35:36',NULL,'2017-08-17 19:35:36'),
	(2,'haohongfan@sina.com','123456','2017-08-17 17:51:28','2017-08-17 18:34:56',NULL,'2017-08-17 18:34:56');

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
