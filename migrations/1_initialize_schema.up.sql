CREATE TABLE `cars` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `owner` varchar(255) DEFAULT NULL,
  `color` varchar(255) DEFAULT NULL,
  `number_of_wheels` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_cars_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- LOCK TABLES `cars` WRITE;

-- UNLOCK TABLES;
