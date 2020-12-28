# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.27)
# Database: listenfield
# Generation Time: 2020-12-28 18:04:58 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table activity
# ------------------------------------------------------------

DROP TABLE IF EXISTS `activity`;

CREATE TABLE `activity` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `farm_id` int(11) NOT NULL,
  `field_id` int(11) NOT NULL,
  `tractor_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `activity_name` enum('PREP','SOWED','FERTILIZED','HARVESTED') DEFAULT NULL,
  `area` double NOT NULL,
  `cost` double NOT NULL,
  `revenue` double NOT NULL,
  `created_date` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `activity` WRITE;
/*!40000 ALTER TABLE `activity` DISABLE KEYS */;

INSERT INTO `activity` (`id`, `farm_id`, `field_id`, `tractor_id`, `user_id`, `activity_name`, `area`, `cost`, `revenue`, `created_date`)
VALUES
	(1,1,3,1,4,'HARVESTED',310,0,31000,'2020-12-28 18:02:10'),
	(2,1,3,1,4,'SOWED',150,3000,0,'2020-12-28 18:02:45');

/*!40000 ALTER TABLE `activity` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table farm
# ------------------------------------------------------------

DROP TABLE IF EXISTS `farm`;

CREATE TABLE `farm` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `farm_name` varchar(255) DEFAULT NULL,
  `farm_owner_id` int(11) NOT NULL,
  `created_date` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `farm` WRITE;
/*!40000 ALTER TABLE `farm` DISABLE KEYS */;

INSERT INTO `farm` (`id`, `farm_name`, `farm_owner_id`, `created_date`)
VALUES
	(1,'jia first farm',2,'2020-12-28 17:44:24');

/*!40000 ALTER TABLE `farm` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table farm_worker
# ------------------------------------------------------------

DROP TABLE IF EXISTS `farm_worker`;

CREATE TABLE `farm_worker` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `farm_id` int(11) unsigned NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `farm_worker` WRITE;
/*!40000 ALTER TABLE `farm_worker` DISABLE KEYS */;

INSERT INTO `farm_worker` (`id`, `farm_id`, `user_id`)
VALUES
	(1,1,3);

/*!40000 ALTER TABLE `farm_worker` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table field
# ------------------------------------------------------------

DROP TABLE IF EXISTS `field`;

CREATE TABLE `field` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `farm_id` int(11) NOT NULL,
  `field_name` varchar(255) NOT NULL,
  `crop` varchar(255) DEFAULT NULL,
  `area` double NOT NULL,
  `created_date` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table tractor
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tractor`;

CREATE TABLE `tractor` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `tractor_name` varchar(255) DEFAULT NULL,
  `farm_id` int(11) NOT NULL,
  `created_date` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `tractor` WRITE;
/*!40000 ALTER TABLE `tractor` DISABLE KEYS */;

INSERT INTO `tractor` (`id`, `tractor_name`, `farm_id`, `created_date`)
VALUES
	(2,'modified jia tractor',0,'2020-12-28 17:53:45');

/*!40000 ALTER TABLE `tractor` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `created_date` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;

INSERT INTO `user` (`id`, `username`, `password`, `email`, `created_date`)
VALUES
	(1,'yanisa','$2a$04$omcGlYdVnzpQXjhWELXMeOOQz339lDyD3zhDZCj1l2/YL.7yo1RYu','chillshyld@gmail.com','2020-12-28 17:28:03'),
	(2,'suradid','$2a$04$A.bgy48kIVyj37rUwLq6XuIlu174jh030So04NHeJEygESUXVr6W2','chillshyld@gmail.com','2020-12-28 17:34:17'),
	(3,'suvinai','$2a$04$ap74hdJWz5nhHHN/CgNnBeSuIdlm5uOEWR9XtVTyRySib9fLgb1f.','suvinai@gmail.com','2020-12-28 17:45:17');

/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
