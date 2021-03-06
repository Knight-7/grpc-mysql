CREATE TABLE `user` (
	`id` BIGINT NOT NULL auto_increment,
	`name` VARCHAR ( 30 ) NOT NULL,
	`age` INT NOT NULL,
	`email` VARCHAR ( 50 ) NOT NULL,
	`create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;