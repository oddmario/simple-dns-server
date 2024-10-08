SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;


CREATE TABLE `dns_records` (
  `id` bigint(20) NOT NULL,
  `record_type` text NOT NULL,
  `record_name` text NOT NULL,
  `record_value` text NOT NULL,
  `record_ttl` bigint(20) NOT NULL DEFAULT 3600,
  `srv_priority` bigint(20) NOT NULL DEFAULT 0,
  `srv_weight` bigint(20) NOT NULL DEFAULT 0,
  `srv_port` bigint(20) NOT NULL DEFAULT 0,
  `srv_target` text NOT NULL,
  `is_disposable` int(11) NOT NULL DEFAULT 0,
  `delete_at_timestamp` bigint(20) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


ALTER TABLE `dns_records`
  ADD PRIMARY KEY (`id`),
  ADD KEY `record_name` (`record_name`(768)),
  ADD KEY `is_disposable` (`is_disposable`),
  ADD KEY `delete_at_timestamp` (`delete_at_timestamp`),
  ADD KEY `record_type` (`record_type`(768));


ALTER TABLE `dns_records`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;