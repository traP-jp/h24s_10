-- +goose Up
CREATE TABLE `applicants` (
    `id` VARCHAR(36) NOT NULL,
    `event_id` VARCHAR(36) NOT NULL,
    `traq_id` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`event_id`) REFERENCES `events` (`id`)
);
