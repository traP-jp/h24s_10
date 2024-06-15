-- +goose Up
CREATE TABLE `events` (
    `id` VARCHAR(36) NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `start` DATETIME,
    `end` DATETIME,
    `host_id` VARCHAR(36) NOT NULL, -- traqのuser id
    `location` VARCHAR(255) NOT NULL,
    `description` TEXT,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `is_confirmed` TINYINT(1) NOT NULL, -- 確定しているかどうか
    PRIMARY KEY (`id`)
);

CREATE TABLE `targets` (
    `id` VARCHAR(36) NOT NULL,
    `event_id` VARCHAR(36) NOT NULL,
    `traq_id` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`event_id`) REFERENCES `events` (`id`)
);

CREATE TABLE `participants` (
    `id` VARCHAR(36) NOT NULL,
    `traq_id` VARCHAR(255) NOT NULL,
    `event_id` VARCHAR(36) NOT NULL,
    `team_id` INT NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`event_id`) REFERENCES `events` (`id`)
);

CREATE TABLE `event_dates` (
    `id` VARCHAR(36) NOT NULL,
    `event_id` VARCHAR(36) NOT NULL,
    `start` DATETIME NOT NULL,
    `end` DATETIME NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`event_id`) REFERENCES `events` (`id`)
);

CREATE TABLE `date_votes` (
    `id` VARCHAR(36) NOT NULL,
    `event_date_id` VARCHAR(36) NOT NULL,
    `user_id` VARCHAR(36) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`event_date_id`) REFERENCES `event_dates` (`id`)
);

CREATE TABLE `comments` (
    `id` VARCHAR(36) NOT NULL,
    `event_id` VARCHAR(36) NOT NULL,
    `user_id` VARCHAR(36) NOT NULL,
    `description` TEXT,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`event_id`) REFERENCES `events` (`id`)
);
