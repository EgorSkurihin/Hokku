USE hokku;

CREATE TABLE `users` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`email` VARCHAR(255) NOT NULL UNIQUE,
	`name` VARCHAR(255) NOT NULL,
	`password` VARCHAR(255) NOT NULL,
    `created` DATETIME NOT NULL, 
	PRIMARY KEY (`id`)
);

CREATE TABLE `themes` (
	`id` INT NOT NULL AUTO_INCREMENT UNIQUE,
	`title` VARCHAR(40) NOT NULL UNIQUE,
	PRIMARY KEY (`id`)
);

CREATE TABLE `hokkus` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(255) NOT NULL, 
	`content` TEXT NOT NULL, 
	`created` DATETIME NOT NULL, 
	`owner` BIGINT NOT NULL,
	`theme` INT NOT NULL,
	PRIMARY KEY (`id`)
);

ALTER TABLE `hokkus` ADD CONSTRAINT `Hokku_fk0` FOREIGN KEY (`owner`) REFERENCES `users`(`id`) ON DELETE CASCADE;

ALTER TABLE `hokkus` ADD CONSTRAINT `Hokku_fk1` FOREIGN KEY (`theme`) REFERENCES `themes`(`id`) ON DELETE CASCADE;

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_themes_id ON themes(id);
CREATE INDEX idx_hokku_id ON hokkus(id);