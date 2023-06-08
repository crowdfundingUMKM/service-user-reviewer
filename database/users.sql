CREATE TABLE `users`
(
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `uReviewer_id` CHAR(12) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255),
    `educational_background` VARCHAR(255),
    `phone` VARCHAR(20),
    `description` TEXT,
    `password_hash` varchar(255) DEFAULT NULL,
    `avatar_file_name` varchar(255) DEFAULT NULL,
    `status_account` varchar(255) DEFAULT NULL,
    `token` varchar(255) DEFAULT NULL,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- data
INSERT INTO `users` (`id`, `uReviewer_id`,`name`, `email`,`educational_background`, `phone`,`description`, `password_hash`, `avatar_file_name`, `status_account`, `token`, `created_at`, `updated_at`) VALUES
(1, '7d4aa4f2-90a', 'Ahmad Zaky', 'test@gmail.com',`Saya merupakan gelar sarjana dari USK `, "82363152828","Ini description", '$2a$04$6A5/psA4hCa0p0mLZQw4A.GKrkYDH3nTiim8lj9mYS18dmVi2FIvO', '', 'active', '', '2023-03-15 22:56:25', '2023-03-15 22:56:25');

-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);
