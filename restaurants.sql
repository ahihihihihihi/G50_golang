/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `restaurants`;
CREATE TABLE `restaurants` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `owner_id` int(11) DEFAULT NULL,
  `name` varchar(50) NOT NULL,
  `addr` varchar(255) NOT NULL,
  `city_id` int(11) DEFAULT NULL,
  `lat` double DEFAULT NULL,
  `lng` double DEFAULT NULL,
  `cover` json DEFAULT NULL,
  `logo` json DEFAULT NULL,
  `type` enum('normal','premium') DEFAULT 'normal',
  `shipping_fee_per_km` double DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `owner_id` (`owner_id`) USING BTREE,
  KEY `city_id` (`city_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

INSERT INTO `restaurants` (`id`, `owner_id`, `name`, `addr`, `city_id`, `lat`, `lng`, `cover`, `logo`, `type`, `shipping_fee_per_km`, `status`, `created_at`, `updated_at`) VALUES
(2, 1, '500lab', '9 Pham Van Hai', NULL, NULL, NULL, NULL, NULL, 'normal', 0, 1, '2023-09-11 16:00:18', '2023-09-12 17:11:06');
INSERT INTO `restaurants` (`id`, `owner_id`, `name`, `addr`, `city_id`, `lat`, `lng`, `cover`, `logo`, `type`, `shipping_fee_per_km`, `status`, `created_at`, `updated_at`) VALUES
(3, 2, 'Tani', '9 Pham Van Hai', NULL, NULL, NULL, NULL, NULL, 'normal', 0, 1, '2023-09-11 16:05:13', '2023-09-12 17:11:06');
INSERT INTO `restaurants` (`id`, `owner_id`, `name`, `addr`, `city_id`, `lat`, `lng`, `cover`, `logo`, `type`, `shipping_fee_per_km`, `status`, `created_at`, `updated_at`) VALUES
(6, NULL, 'Tani', '9 Pham Van Hai', NULL, NULL, NULL, NULL, NULL, 'normal', 0, 1, '2023-09-12 03:48:05', '2023-09-12 03:48:05');
INSERT INTO `restaurants` (`id`, `owner_id`, `name`, `addr`, `city_id`, `lat`, `lng`, `cover`, `logo`, `type`, `shipping_fee_per_km`, `status`, `created_at`, `updated_at`) VALUES
(7, NULL, 'update name restaurant', '2 Le Duan', NULL, NULL, NULL, NULL, NULL, 'normal', 0, 1, '2023-09-12 04:33:23', '2023-09-12 08:04:00'),
(8, NULL, 'a new restaurant 2', 'somewhere 2', NULL, NULL, NULL, NULL, NULL, 'normal', 0, 1, '2023-09-12 10:19:24', '2023-09-12 10:19:24'),
(9, NULL, 'a new restaurant 3', 'somewhere 3', NULL, NULL, NULL, NULL, NULL, 'normal', 0, 0, '2023-09-12 10:19:37', '2023-09-12 14:59:03'),
(10, NULL, 'a new restaurant rant', 'somewhere 5', NULL, NULL, NULL, NULL, NULL, 'normal', 0, 1, '2023-09-13 10:24:23', '2023-09-13 10:24:23'),
(11, NULL, 'a new restaurant rant 6', 'somewhere 6', NULL, NULL, NULL, NULL, NULL, 'normal', 0, 1, '2023-09-13 11:07:04', '2023-09-13 11:07:04');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;