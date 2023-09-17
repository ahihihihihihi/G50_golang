/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `carts`;
CREATE TABLE `carts` (
  `user_id` int(11) NOT NULL,
  `food_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`food_id`),
  KEY `food_id` (`food_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` text,
  `icon` json DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `cities`;
CREATE TABLE `cities` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 NOT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `food_likes`;
CREATE TABLE `food_likes` (
  `user_id` int(11) NOT NULL,
  `food_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`food_id`),
  KEY `food_id` (`food_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `food_ratings`;
CREATE TABLE `food_ratings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `food_id` int(11) NOT NULL,
  `point` float DEFAULT '0',
  `comment` text,
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `food_id` (`food_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `foods`;
CREATE TABLE `foods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `restaurant_id` int(11) NOT NULL,
  `category_id` int(11) DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `price` float NOT NULL,
  `images` json NOT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `restaurant_id` (`restaurant_id`) USING BTREE,
  KEY `category_id` (`category_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `order_details`;
CREATE TABLE `order_details` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) NOT NULL,
  `food_origin` json DEFAULT NULL,
  `price` float NOT NULL,
  `quantity` int(11) NOT NULL,
  `discount` float DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `order_trackings`;
CREATE TABLE `order_trackings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) NOT NULL,
  `state` enum('waiting_for_shipper','preparing','on_the_way','delivered','cancel') NOT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `total_price` float NOT NULL,
  `shipper_id` int(11) DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `shipper_id` (`shipper_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `restaurant_foods`;
CREATE TABLE `restaurant_foods` (
  `restaurant_id` int(11) NOT NULL,
  `food_id` int(11) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`restaurant_id`,`food_id`),
  KEY `food_id` (`food_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `restaurant_likes`;
CREATE TABLE `restaurant_likes` (
  `restaurant_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`restaurant_id`,`user_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `restaurant_ratings`;
CREATE TABLE `restaurant_ratings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `restaurant_id` int(11) NOT NULL,
  `point` float NOT NULL DEFAULT '0',
  `comment` text,
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `restaurant_id` (`restaurant_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `restaurants`;
CREATE TABLE `restaurants` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `owner_id` int(11) DEFAULT NULL,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `addr` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `city_id` int(11) DEFAULT NULL,
  `lat` double DEFAULT NULL,
  `lng` double DEFAULT NULL,
  `cover` json DEFAULT NULL,
  `logo` json DEFAULT NULL,
  `shipping_fee_per_km` double DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `owner_id` (`owner_id`) USING BTREE,
  KEY `city_id` (`city_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `user_addresses`;
CREATE TABLE `user_addresses` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `city_id` int(11) NOT NULL,
  `title` varchar(100) DEFAULT NULL,
  `icon` json DEFAULT NULL,
  `addr` varchar(255) NOT NULL,
  `lat` double DEFAULT NULL,
  `lng` double DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `city_id` (`city_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `user_device_tokens`;
CREATE TABLE `user_device_tokens` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned DEFAULT NULL,
  `is_production` tinyint(1) DEFAULT '0',
  `os` enum('ios','android','web') DEFAULT 'ios' COMMENT '1: iOS, 2: Android',
  `token` varchar(255) DEFAULT NULL,
  `status` smallint(5) unsigned NOT NULL DEFAULT '1',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `os` (`os`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `fb_id` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `gg_id` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `salt` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `last_name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `first_name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `role` enum('user','admin','shipper') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'user',
  `avatar` json DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;





INSERT INTO `cities` (`id`, `title`, `status`, `created_at`, `updated_at`) VALUES
(1, 'An Giang', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22');
INSERT INTO `cities` (`id`, `title`, `status`, `created_at`, `updated_at`) VALUES
(2, 'Vũng Tàu', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22');
INSERT INTO `cities` (`id`, `title`, `status`, `created_at`, `updated_at`) VALUES
(3, 'Bắc Giang', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22');
INSERT INTO `cities` (`id`, `title`, `status`, `created_at`, `updated_at`) VALUES
(4, 'Bắc Cạn', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(5, 'Bạc Liêu', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(6, 'Bắc Ninh', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(7, 'Bến Tre', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(8, 'Bình Định', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(9, 'Bình Dương', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(10, 'Bình Phước', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(11, 'Bình Thuận', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(12, 'Cà Mau', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(13, 'Cần Thơ', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(14, 'Cao Bằng', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(15, 'Đà Nẵng', 1, '2020-05-18 06:52:21', '2020-05-18 06:52:21'),
(16, 'Đắk Lắk', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(17, 'Đắk Nông', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(18, 'Điện Biên', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(19, 'Đồng Nai', 1, '2020-05-18 06:52:21', '2020-05-18 06:52:21'),
(20, 'Đồng Tháp', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(21, 'Gia Lai', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(22, 'Hà Giang', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(23, 'Hà Nam', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(24, 'Hà Nội', 1, '2020-05-18 06:52:21', '2020-05-18 06:52:21'),
(25, 'Hà Tĩnh', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(26, 'Hải Dương', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(27, 'Hải Phòng', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(28, 'Hậu Giang', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(29, 'Hoà Bình', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(30, 'Hưng Yên', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(31, 'Khánh Hoà', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(32, 'Kiên Giang', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(33, 'Kon Tum', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(34, 'Lai Châu', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(35, 'Lâm Đồng', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(36, 'Lạng Sơn', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(37, 'Lào Cai', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(38, 'Long An', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(39, 'Nam Định', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(40, 'Nghệ An', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(41, 'Ninh Bình', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(42, 'Ninh Thuận', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(43, 'Phú Thọ', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(44, 'Phú Yên', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(45, 'Quảng Bình', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(46, 'Quảng Namm', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(47, 'Quãng Ngãi', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(48, 'Quãng Ninh', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(49, 'Quãng Trị', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(50, 'Sóc Trăng', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(51, 'Sơn La', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(52, 'Tây Ninh', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(53, 'Thái Bình', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(54, 'Thái Nguyên', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(55, 'Thanh Hoá', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(56, 'Huế', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(57, 'Tiền Giang', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(58, 'Hồ Chí Minh', 1, '2020-05-18 06:41:51', '2020-05-18 06:41:51'),
(59, 'Trà Vinh', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(60, 'Tuyên Quang', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(61, 'Vĩnh Long', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(62, 'Vĩnh Phúc', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(63, 'Yên Bái', 1, '2020-05-18 06:55:19', '2020-05-18 06:55:19');



















INSERT INTO `restaurants` (`id`, `owner_id`, `name`, `addr`, `city_id`, `lat`, `lng`, `cover`, `logo`, `shipping_fee_per_km`, `status`, `created_at`, `updated_at`) VALUES
(1, NULL, 'Name 1', 'Address 1', NULL, NULL, NULL, NULL, NULL, 0, 1, '2023-09-14 11:13:52', '2023-09-14 11:13:52');
INSERT INTO `restaurants` (`id`, `owner_id`, `name`, `addr`, `city_id`, `lat`, `lng`, `cover`, `logo`, `shipping_fee_per_km`, `status`, `created_at`, `updated_at`) VALUES
(2, NULL, 'Name 2', 'Address 2', NULL, NULL, NULL, NULL, NULL, 0, 1, '2023-09-14 11:13:52', '2023-09-14 11:13:52');
INSERT INTO `restaurants` (`id`, `owner_id`, `name`, `addr`, `city_id`, `lat`, `lng`, `cover`, `logo`, `shipping_fee_per_km`, `status`, `created_at`, `updated_at`) VALUES
(3, NULL, 'Name 3', 'Address 3', NULL, NULL, NULL, NULL, NULL, 0, 1, '2023-09-14 11:13:52', '2023-09-14 11:13:52');
INSERT INTO `restaurants` (`id`, `owner_id`, `name`, `addr`, `city_id`, `lat`, `lng`, `cover`, `logo`, `shipping_fee_per_km`, `status`, `created_at`, `updated_at`) VALUES
(4, NULL, 'Name 4', 'Address 4', NULL, NULL, NULL, NULL, NULL, 0, 1, '2023-09-14 11:13:52', '2023-09-14 11:13:52'),
(5, NULL, 'Name 5', 'Address 5', NULL, NULL, NULL, NULL, NULL, 0, 1, '2023-09-14 11:13:52', '2023-09-14 11:13:52'),
(6, NULL, 'a new restaurant rant 7', 'somewhere 7', NULL, NULL, NULL, NULL, NULL, 0, 0, '2023-09-15 11:38:18', '2023-09-15 04:43:52'),
(7, NULL, 'a new restaurant rant 8', 'somewhere 8', NULL, NULL, NULL, NULL, '{\"id\": 0, \"url\": \"https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/Pizza_Hut_classic_logo.svg/800px-Pizza_Hut_classic_logo.svg.png\", \"width\": 800, \"height\": 642}', 0, 1, '2023-09-15 16:45:54', '2023-09-15 16:45:54'),
(8, NULL, 'a new restaurant rant 9', 'somewhere 9', NULL, NULL, NULL, '[{\"id\": 0, \"url\": \"https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/Pizza_Hut_classic_logo.svg/800px-Pizza_Hut_classic_logo.svg.png\", \"width\": 800, \"height\": 642}, {\"id\": 0, \"url\": \"https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/Pizza_Hut_classic_logo.svg/800px-Pizza_Hut_classic_logo.svg.png\", \"width\": 800, \"height\": 642}]', '{\"id\": 0, \"url\": \"https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/Pizza_Hut_classic_logo.svg/800px-Pizza_Hut_classic_logo.svg.png\", \"width\": 800, \"height\": 642}', 0, 1, '2023-09-15 17:04:44', '2023-09-15 17:04:44');





INSERT INTO `users` (`id`, `email`, `fb_id`, `gg_id`, `password`, `salt`, `last_name`, `first_name`, `phone`, `role`, `avatar`, `status`, `created_at`, `updated_at`) VALUES
(4, 'du@gmail.com', NULL, NULL, '33be1340e39a2c6dc8e26937278d90a2', 'QIXIFUHeHAFqXdIEXrTKsYXeKOYpBwHRqOWVjkVqYXsAHVXZwc', 'Con', 'Du', '', 'user', NULL, 1, '2023-09-17 16:05:02', '2023-09-17 16:05:02');



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;