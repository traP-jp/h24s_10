-- +goose Up
ALTER TABLE `date_votes` RENAME COLUMN `user_id` TO `traq_id`;
ALTER TABLE `comments` RENAME COLUMN `user_id` TO `traq_id`;
ALTER TABLE `comments` RENAME COLUMN `description` TO `content`;
