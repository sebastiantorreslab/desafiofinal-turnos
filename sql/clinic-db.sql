-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema clinic-db
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema clinic-db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `clinic-db` DEFAULT CHARACTER SET latin1 ;
USE `clinic-db` ;

-- -----------------------------------------------------
-- Table `clinic-db`.`dentists`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clinic-db`.`dentists` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `license` VARCHAR(45) NULL DEFAULT NULL,
  `name` VARCHAR(45) NULL DEFAULT NULL,
  `last_name` VARCHAR(45) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 3
DEFAULT CHARACTER SET = latin1;


-- -----------------------------------------------------
-- Table `clinic-db`.`patients`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clinic-db`.`patients` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `dni` INT(11) NULL DEFAULT NULL,
  `name` VARCHAR(45) NULL DEFAULT NULL,
  `last_name` VARCHAR(45) NULL DEFAULT NULL,
  `address` VARCHAR(45) NULL DEFAULT NULL,
  `admission_date` VARCHAR(45) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = latin1;


-- -----------------------------------------------------
-- Table `clinic-db`.`shift`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clinic-db`.`shift` (
  `id` INT(11) NOT NULL,
  `shift_date` VARCHAR(45) NOT NULL,
  `shift_hour` VARCHAR(45) NOT NULL,
  `id_patient` INT(11) NULL DEFAULT NULL,
  `id_dentist` INT(11) NULL DEFAULT NULL,
  INDEX `id_patient_idx` (`id_patient` ASC) VISIBLE,
  INDEX `id_dentist_idx` (`id_dentist` ASC) VISIBLE,
  CONSTRAINT `id_dentist`
    FOREIGN KEY (`id_dentist`)
    REFERENCES `clinic-db`.`dentists` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `id_patient`
    FOREIGN KEY (`id_patient`)
    REFERENCES `clinic-db`.`patients` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = latin1;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
