# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: localhost (MySQL 5.7.15)
# Database: devtest
# Generation Time: 2016-10-03 09:24:48 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table AUTH
# ------------------------------------------------------------

DROP TABLE IF EXISTS `AUTH`;

CREATE TABLE `AUTH` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `EMAIL` varchar(128) NOT NULL DEFAULT '',
  `OTP` int(6) unsigned NOT NULL,
  `ATTEMPTS` tinyint(1) DEFAULT NULL,
  `ACTIVE` tinyint(1) NOT NULL DEFAULT '0',
  `EXPIRY` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_email` (`EMAIL`),
  KEY `idx_otp` (`OTP`),
  KEY `idx_email` (`EMAIL`),
  KEY `idx_active` (`ACTIVE`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `AUTH` WRITE;
/*!40000 ALTER TABLE `AUTH` DISABLE KEYS */;

INSERT INTO `AUTH` (`id`, `EMAIL`, `OTP`, `ATTEMPTS`, `ACTIVE`, `EXPIRY`)
VALUES
	(1,'deepak@skidata.com',2272,0,1,'2016-10-04 05:55:26');

/*!40000 ALTER TABLE `AUTH` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table DOMAIN
# ------------------------------------------------------------

DROP TABLE IF EXISTS `DOMAIN`;

CREATE TABLE `DOMAIN` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `DOMAIN` varchar(128) DEFAULT NULL,
  `ALLOW` tinyint(1) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_domain` (`DOMAIN`),
  KEY `idx_domain` (`DOMAIN`),
  KEY `idx_allow` (`ALLOW`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `DOMAIN` WRITE;
/*!40000 ALTER TABLE `DOMAIN` DISABLE KEYS */;

INSERT INTO `DOMAIN` (`id`, `DOMAIN`, `ALLOW`)
VALUES
	(1,'infosys.com',1),
	(2,'skidata.com',1),
	(3,'ansrsource.com',1),
	(4,'hotmail.com',0),
	(5,'gmail.com',0),
	(6,'yahoo.com',0);

/*!40000 ALTER TABLE `DOMAIN` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table USER
# ------------------------------------------------------------

DROP TABLE IF EXISTS `USER`;

CREATE TABLE `USER` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `EMAIL` varchar(128) DEFAULT NULL,
  `FIRSTNAME` varchar(64) DEFAULT NULL,
  `LASTNAME` varchar(64) DEFAULT NULL,
  `COUNTRYCODE` tinyint(4) unsigned DEFAULT NULL,
  `PHONE` bigint(16) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_email` (`EMAIL`),
  KEY `idx_email` (`EMAIL`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
