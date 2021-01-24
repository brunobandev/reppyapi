CREATE TABLE `users` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`email` VARCHAR(255) NOT NULL UNIQUE,
	`password` VARCHAR(255) NOT NULL,
	`createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
);

CREATE TABLE `properties` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`userId` INT(11) NOT NULL,
	`name` varchar(255) NOT NULL,
	`price` FLOAT(7,2) NOT NULL,
	`createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
);

CREATE TABLE `accommodations` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`propertyId` INT(11) NOT NULL,
	`name` VARCHAR(255) NOT NULL,
	`price` FLOAT(7,2) NOT NULL,
	`createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
);

CREATE TABLE `accommodation_user` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`accommodationId` INT(11) NOT NULL,
	`userId` INT(11) NOT NULL,
	`active` BOOLEAN NOT NULL DEFAULT '1',
	`createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
);

CREATE TABLE `bills` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`propertyId` INT(11) NOT NULL,
	`description` VARCHAR(255) NOT NULL,
	`price` FLOAT(7,2) NOT NULL,
	`file` VARCHAR(255),
	`dueAt` DATETIME NOT NULL,
	`createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
);

CREATE TABLE `bill_user` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`billId` INT(11) NOT NULL,
	`userId` INT(11) NOT NULL,
	`price` FLOAT(7,2) NOT NULL,
	`paid` BOOLEAN NOT NULL DEFAULT '0',
	`createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
);

ALTER TABLE `properties` ADD CONSTRAINT `properties_fk0` FOREIGN KEY (`userId`) REFERENCES `users`(`id`);

ALTER TABLE `accommodations` ADD CONSTRAINT `accommodations_fk0` FOREIGN KEY (`propertyId`) REFERENCES `properties`(`id`);

ALTER TABLE `accommodation_user` ADD CONSTRAINT `accommodation_user_fk0` FOREIGN KEY (`accommodationId`) REFERENCES `accommodations`(`id`);

ALTER TABLE `accommodation_user` ADD CONSTRAINT `accommodation_user_fk1` FOREIGN KEY (`userId`) REFERENCES `users`(`id`);

ALTER TABLE `bills` ADD CONSTRAINT `bills_fk0` FOREIGN KEY (`propertyId`) REFERENCES `properties`(`id`);

ALTER TABLE `bill_user` ADD CONSTRAINT `bill_user_fk0` FOREIGN KEY (`billId`) REFERENCES `bills`(`id`);

ALTER TABLE `bill_user` ADD CONSTRAINT `bill_user_fk1` FOREIGN KEY (`userId`) REFERENCES `users`(`id`);