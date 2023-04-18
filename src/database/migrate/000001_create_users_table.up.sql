CREATE TABLE IF NOT EXISTS `users` (
  `id`         int(11) NOT NULL AUTO_INCREMENT,
  `uid`        VARCHAR(191) NOT NULL,
  `name`       VARCHAR(191) NULL,
  `email`      VARCHAR(191) NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL,
  UNIQUE INDEX `users_email_key`(`email`),
  UNIQUE INDEX `users_uid_key`(`uid`),
  PRIMARY KEY (`id`)
);