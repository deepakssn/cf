SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL';

CREATE SCHEMA IF NOT EXISTS `try` DEFAULT CHARACTER SET latin1 ;
USE `try` ;

-- -----------------------------------------------------
-- Table `try`.`AUTH`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `try`.`AUTH` ;

CREATE  TABLE IF NOT EXISTS `try`.`AUTH` (
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT ,
  `EMAIL` VARCHAR(128) NOT NULL DEFAULT '' ,
  `OTP` INT(6) UNSIGNED NOT NULL ,
  `ATTEMPTS` TINYINT(1) NULL DEFAULT NULL ,
  `IP` VARCHAR(64) NULL DEFAULT NULL ,
  `ACTIVE` TINYINT(1) NOT NULL DEFAULT '0' ,
  PRIMARY KEY (`id`) ,
  INDEX `idx_otp` (`OTP` ASC) ,
  INDEX `idx_email` (`EMAIL` ASC) ,
  INDEX `idx_active` (`ACTIVE` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `try`.`Location`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `try`.`Location` ;

CREATE  TABLE IF NOT EXISTS `try`.`Location` (
  `locId` INT(11) NOT NULL AUTO_INCREMENT ,
  `city` VARCHAR(50) NOT NULL ,
  `addressLine1` VARCHAR(50) NOT NULL ,
  `state` VARCHAR(50) NULL DEFAULT NULL ,
  `country` VARCHAR(50) NOT NULL ,
  `postalCode` VARCHAR(15) NOT NULL ,
  `territory` VARCHAR(10) NOT NULL ,
  PRIMARY KEY (`locId`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `try`.`User`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `try`.`User` ;

CREATE  TABLE IF NOT EXISTS `try`.`User` (
  `userId` INT(11) NOT NULL AUTO_INCREMENT ,
  `userName` VARCHAR(50) NOT NULL ,
  `userSurname` VARCHAR(50) NULL DEFAULT NULL ,
  `userFirstname` VARCHAR(50) NULL DEFAULT NULL ,
  `phone` VARCHAR(50) NOT NULL ,
  `userLoc` INT(11) NOT NULL ,
  `email` VARCHAR(128) NOT NULL ,
  PRIMARY KEY (`userId`) ,
  INDEX `userLoc_fk` (`userLoc` ASC) ,
  CONSTRAINT `userLoc_fk`
    FOREIGN KEY (`userLoc` )
    REFERENCES `try`.`Location` (`locId` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `try`.`Blip`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `try`.`Blip` ;

CREATE  TABLE IF NOT EXISTS `try`.`Blip` (
  `blipId` INT(11) NOT NULL AUTO_INCREMENT ,
  `postuser` INT(11) NOT NULL ,
  PRIMARY KEY (`blipId`) ,
  INDEX `postuser_fk` (`postuser` ASC) ,
  CONSTRAINT `postuser_fk`
    FOREIGN KEY (`postuser` )
    REFERENCES `try`.`User` (`userId` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `try`.`DOMAIN`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `try`.`DOMAIN` ;

CREATE  TABLE IF NOT EXISTS `try`.`DOMAIN` (
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT ,
  `DOMAIN` VARCHAR(128) NOT NULL ,
  `ALLOW` TINYINT(1) UNSIGNED NULL DEFAULT '0' ,
  PRIMARY KEY (`id`) ,
  UNIQUE INDEX `unique_domain` (`DOMAIN` ASC) ,
  INDEX `idx_domain` (`DOMAIN` ASC) ,
  INDEX `idx_allow` (`ALLOW` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;



SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
