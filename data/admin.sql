-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.7.26 - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  12.4.0.6659
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 dcat-admin 的数据库结构
CREATE DATABASE IF NOT EXISTS `dcat-admin` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `dcat-admin`;

-- 导出  表 dcat-admin.admin_extensions 结构
CREATE TABLE IF NOT EXISTS `admin_extensions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `version` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `is_enabled` tinyint(4) NOT NULL DEFAULT '0',
  `options` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_extensions_name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.admin_extensions 的数据：~0 rows (大约)

-- 导出  表 dcat-admin.admin_extension_histories 结构
CREATE TABLE IF NOT EXISTS `admin_extension_histories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `type` tinyint(4) NOT NULL DEFAULT '1',
  `version` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `detail` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `admin_extension_histories_name_index` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.admin_extension_histories 的数据：~0 rows (大约)

-- 导出  表 dcat-admin.admin_menu 结构
CREATE TABLE IF NOT EXISTS `admin_menu` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) NOT NULL DEFAULT '0',
  `order` int(11) NOT NULL DEFAULT '0',
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `uri` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `extension` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `show` tinyint(4) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.admin_menu 的数据：10 rows
/*!40000 ALTER TABLE `admin_menu` DISABLE KEYS */;
REPLACE INTO `admin_menu` (`id`, `parent_id`, `order`, `title`, `icon`, `uri`, `extension`, `show`, `created_at`, `updated_at`) VALUES
	(1, 0, 1, 'Index', 'feather icon-bar-chart-2', '/', '', 1, '2024-08-12 10:42:22', NULL),
	(2, 0, 2, 'Admin', 'feather icon-settings', '', '', 1, '2024-08-12 10:42:22', NULL),
	(3, 2, 3, 'Users', '', 'auth/users', '', 1, '2024-08-12 10:42:22', NULL),
	(4, 2, 4, 'Roles', '', 'auth/roles', '', 1, '2024-08-12 10:42:22', NULL),
	(5, 2, 5, 'Permission', '', 'auth/permissions', '', 1, '2024-08-12 10:42:22', NULL),
	(6, 2, 6, 'Menu', '', 'auth/menu', '', 1, '2024-08-12 10:42:22', NULL),
	(7, 2, 7, 'Extensions', '', 'auth/extensions', '', 1, '2024-08-12 10:42:22', NULL),
	(8, 0, 8, '代理服务器节点列表', 'fa-automobile', '/proxies', '', 1, '2024-08-14 08:50:33', '2024-08-14 09:07:09'),
	(9, 0, 9, '原始订阅地址', 'fa-amazon', 'urls', '', 1, '2024-08-14 08:59:19', '2024-08-14 09:07:25'),
	(13, 0, 10, '子订阅', NULL, 'subInfo', '', 1, '2024-08-16 08:49:55', '2024-08-16 08:49:55');
/*!40000 ALTER TABLE `admin_menu` ENABLE KEYS */;

-- 导出  表 dcat-admin.admin_permissions 结构
CREATE TABLE IF NOT EXISTS `admin_permissions` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `http_method` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `http_path` text COLLATE utf8mb4_unicode_ci,
  `order` int(11) NOT NULL DEFAULT '0',
  `parent_id` bigint(20) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_permissions_slug_unique` (`slug`)
) ENGINE=MyISAM AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.admin_permissions 的数据：6 rows
/*!40000 ALTER TABLE `admin_permissions` DISABLE KEYS */;
REPLACE INTO `admin_permissions` (`id`, `name`, `slug`, `http_method`, `http_path`, `order`, `parent_id`, `created_at`, `updated_at`) VALUES
	(1, 'Auth management', 'auth-management', '', '', 1, 0, '2024-08-12 10:42:22', NULL),
	(2, 'Users', 'users', '', '/auth/users*', 2, 1, '2024-08-12 10:42:22', NULL),
	(3, 'Roles', 'roles', '', '/auth/roles*', 3, 1, '2024-08-12 10:42:22', NULL),
	(4, 'Permissions', 'permissions', '', '/auth/permissions*', 4, 1, '2024-08-12 10:42:22', NULL),
	(5, 'Menu', 'menu', '', '/auth/menu*', 5, 1, '2024-08-12 10:42:22', NULL),
	(6, 'Extension', 'extension', '', '/auth/extensions*', 6, 1, '2024-08-12 10:42:22', NULL);
/*!40000 ALTER TABLE `admin_permissions` ENABLE KEYS */;

-- 导出  表 dcat-admin.admin_permission_menu 结构
CREATE TABLE IF NOT EXISTS `admin_permission_menu` (
  `permission_id` bigint(20) NOT NULL,
  `menu_id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  UNIQUE KEY `admin_permission_menu_permission_id_menu_id_unique` (`permission_id`,`menu_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.admin_permission_menu 的数据：12 rows
/*!40000 ALTER TABLE `admin_permission_menu` DISABLE KEYS */;
REPLACE INTO `admin_permission_menu` (`permission_id`, `menu_id`, `created_at`, `updated_at`) VALUES
	(1, 8, '2024-08-14 08:50:33', '2024-08-14 08:50:33'),
	(2, 8, '2024-08-14 08:50:33', '2024-08-14 08:50:33'),
	(3, 8, '2024-08-14 08:50:33', '2024-08-14 08:50:33'),
	(4, 8, '2024-08-14 08:50:33', '2024-08-14 08:50:33'),
	(5, 8, '2024-08-14 08:50:33', '2024-08-14 08:50:33'),
	(6, 8, '2024-08-14 08:50:33', '2024-08-14 08:50:33'),
	(1, 9, '2024-08-14 08:59:19', '2024-08-14 08:59:19'),
	(2, 9, '2024-08-14 08:59:19', '2024-08-14 08:59:19'),
	(3, 9, '2024-08-14 08:59:19', '2024-08-14 08:59:19'),
	(4, 9, '2024-08-14 08:59:19', '2024-08-14 08:59:19'),
	(5, 9, '2024-08-14 08:59:19', '2024-08-14 08:59:19'),
	(6, 9, '2024-08-14 08:59:19', '2024-08-14 08:59:19');
/*!40000 ALTER TABLE `admin_permission_menu` ENABLE KEYS */;

-- 导出  表 dcat-admin.admin_roles 结构
CREATE TABLE IF NOT EXISTS `admin_roles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_roles_slug_unique` (`slug`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.admin_roles 的数据：1 rows
/*!40000 ALTER TABLE `admin_roles` DISABLE KEYS */;
REPLACE INTO `admin_roles` (`id`, `name`, `slug`, `created_at`, `updated_at`) VALUES
	(1, 'Administrator', 'administrator', '2024-08-12 10:42:22', '2024-08-12 10:42:22');
/*!40000 ALTER TABLE `admin_roles` ENABLE KEYS */;

-- 导出  表 dcat-admin.admin_role_menu 结构
CREATE TABLE IF NOT EXISTS `admin_role_menu` (
  `role_id` bigint(20) NOT NULL,
  `menu_id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  UNIQUE KEY `admin_role_menu_role_id_menu_id_unique` (`role_id`,`menu_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.admin_role_menu 的数据：3 rows
/*!40000 ALTER TABLE `admin_role_menu` DISABLE KEYS */;
REPLACE INTO `admin_role_menu` (`role_id`, `menu_id`, `created_at`, `updated_at`) VALUES
	(1, 8, '2024-08-14 08:50:33', '2024-08-14 08:50:33'),
	(1, 9, '2024-08-14 08:59:19', '2024-08-14 08:59:19'),
	(1, 13, '2024-08-16 08:49:55', '2024-08-16 08:49:55');
/*!40000 ALTER TABLE `admin_role_menu` ENABLE KEYS */;

-- 导出  表 dcat-admin.admin_role_permissions 结构
CREATE TABLE IF NOT EXISTS `admin_role_permissions` (
  `role_id` bigint(20) NOT NULL,
  `permission_id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  UNIQUE KEY `admin_role_permissions_role_id_permission_id_unique` (`role_id`,`permission_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.admin_role_permissions 的数据：0 rows
/*!40000 ALTER TABLE `admin_role_permissions` DISABLE KEYS */;
/*!40000 ALTER TABLE `admin_role_permissions` ENABLE KEYS */;

-- 导出  表 dcat-admin.admin_role_users 结构
CREATE TABLE IF NOT EXISTS `admin_role_users` (
  `role_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  UNIQUE KEY `admin_role_users_role_id_user_id_unique` (`role_id`,`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.admin_role_users 的数据：1 rows
/*!40000 ALTER TABLE `admin_role_users` DISABLE KEYS */;
REPLACE INTO `admin_role_users` (`role_id`, `user_id`, `created_at`, `updated_at`) VALUES
	(1, 1, '2024-08-12 10:42:22', '2024-08-12 10:42:22');
/*!40000 ALTER TABLE `admin_role_users` ENABLE KEYS */;

-- 导出  表 dcat-admin.admin_settings 结构
CREATE TABLE IF NOT EXISTS `admin_settings` (
  `slug` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`slug`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.admin_settings 的数据：0 rows
/*!40000 ALTER TABLE `admin_settings` DISABLE KEYS */;
/*!40000 ALTER TABLE `admin_settings` ENABLE KEYS */;

-- 导出  表 dcat-admin.admin_users 结构
CREATE TABLE IF NOT EXISTS `admin_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(120) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(80) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_users_username_unique` (`username`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.admin_users 的数据：1 rows
/*!40000 ALTER TABLE `admin_users` DISABLE KEYS */;
REPLACE INTO `admin_users` (`id`, `username`, `password`, `name`, `avatar`, `remember_token`, `created_at`, `updated_at`) VALUES
	(1, 'admin', '$2y$10$2.9wOt3Eq9KDh6NToaScc.FfXhENw9qSPpR59REUpPO/CjQi71MRi', 'Administrator', NULL, 'UeF7XSbA7Q7GFuBINqU3wfarNYAMOxI4Xr4Y4tni9rx2632AXhXYPgDYqZ3K', '2024-08-12 10:42:22', '2024-08-12 10:42:22');
/*!40000 ALTER TABLE `admin_users` ENABLE KEYS */;

-- 导出  表 dcat-admin.failed_jobs 结构
CREATE TABLE IF NOT EXISTS `failed_jobs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `connection` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `queue` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `payload` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `exception` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `failed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `failed_jobs_uuid_unique` (`uuid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.failed_jobs 的数据：0 rows
/*!40000 ALTER TABLE `failed_jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `failed_jobs` ENABLE KEYS */;

-- 导出  表 dcat-admin.migrations 结构
CREATE TABLE IF NOT EXISTS `migrations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.migrations 的数据：8 rows
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;
REPLACE INTO `migrations` (`id`, `migration`, `batch`) VALUES
	(1, '2014_10_12_000000_create_users_table', 1),
	(2, '2014_10_12_100000_create_password_resets_table', 1),
	(3, '2016_01_04_173148_create_admin_tables', 1),
	(4, '2019_08_19_000000_create_failed_jobs_table', 1),
	(5, '2019_12_14_000001_create_personal_access_tokens_table', 1),
	(6, '2020_09_07_090635_create_admin_settings_table', 1),
	(7, '2020_09_22_015815_create_admin_extensions_table', 1),
	(8, '2020_11_01_083237_update_admin_menu_table', 1);
/*!40000 ALTER TABLE `migrations` ENABLE KEYS */;

-- 导出  表 dcat-admin.password_resets 结构
CREATE TABLE IF NOT EXISTS `password_resets` (
  `email` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`email`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.password_resets 的数据：0 rows
/*!40000 ALTER TABLE `password_resets` DISABLE KEYS */;
/*!40000 ALTER TABLE `password_resets` ENABLE KEYS */;

-- 导出  表 dcat-admin.personal_access_tokens 结构
CREATE TABLE IF NOT EXISTS `personal_access_tokens` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `tokenable_type` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tokenable_id` bigint(20) unsigned NOT NULL,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `abilities` text COLLATE utf8mb4_unicode_ci,
  `last_used_at` timestamp NULL DEFAULT NULL,
  `expires_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `personal_access_tokens_token_unique` (`token`),
  KEY `personal_access_tokens_tokenable_type_tokenable_id_index` (`tokenable_type`,`tokenable_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.personal_access_tokens 的数据：0 rows
/*!40000 ALTER TABLE `personal_access_tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `personal_access_tokens` ENABLE KEYS */;

-- 导出  表 dcat-admin.proxies 结构
CREATE TABLE IF NOT EXISTS `proxies` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '节点名称',
  `server` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '节点地址',
  `conf` text COLLATE utf8_unicode_ci NOT NULL COMMENT '详细配置信息',
  `proxie_type` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT 'user' COMMENT '节点类型（自动添加，手动添加）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=46 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- 正在导出表  dcat-admin.proxies 的数据：22 rows
/*!40000 ALTER TABLE `proxies` DISABLE KEYS */;
REPLACE INTO `proxies` (`id`, `name`, `server`, `conf`, `proxie_type`) VALUES
	(44, 'Bronze-美国-onep-02', 'zz.mgxyt.fogvip-zz.uk', '{"name":"Bronze-美国-onep-02","type":"trojan","server":"zz.mgxyt.fogvip-zz.uk","port":40006,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.mgxyt.yuiuingkig.icu"}', 'auto'),
	(43, 'Bronze-新加坡-onep-01', 'zz.xinjp01.fogvip-zz.uk', '{"name":"Bronze-新加坡-onep-01","type":"trojan","server":"zz.xinjp01.fogvip-zz.uk","port":40013,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.xinjp01.yuiuingkig.icu"}', 'auto'),
	(42, 'Bronze-香港-BGP-01', 'zz.xg01.fogvip-zz.uk', '{"name":"Bronze-香港-BGP-01","type":"trojan","server":"zz.xg01.fogvip-zz.uk","port":40001,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.xgg01.yuiuingkig.icu"}', 'auto'),
	(41, 'Silver-香港-HKG-02', 'zz.hgkneko.viptap-tcb-zz.cc', '{"name":"Silver-香港-HKG-02","type":"trojan","server":"zz.hgkneko.viptap-tcb-zz.cc","port":50112,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.hkgneko.yuiuingkig.icu"}', 'auto'),
	(40, 'Silver-日本-商宽-02', 'zz.jp03.fogvip-zz.uk', '{"name":"Silver-日本-商宽-02","type":"trojan","server":"zz.jp03.fogvip-zz.uk","port":50056,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.rbv2.yuiuingkig.icu"}', 'auto'),
	(39, 'Silver-美国-Vegas-04', 'zz.ysls.fogvip-zz.uk', '{"name":"Silver-美国-Vegas-04","type":"trojan","server":"zz.ysls.fogvip-zz.uk","port":50308,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.vegas.yuiuingkig.icu"}', 'auto'),
	(38, 'Bronze-香港-Lite-02', 'zz.xgzl.fogvip-zz.uk', '{"name":"Bronze-香港-Lite-02","type":"trojan","server":"zz.xgzl.fogvip-zz.uk","port":50069,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.xgzl02.yuiuingkig.icu"}', 'auto'),
	(37, 'Silver-美国-Host-01', 'zz.usa01.fogvip-zz.uk', '{"name":"Silver-美国-Host-01","type":"trojan","server":"zz.usa01.fogvip-zz.uk","port":50072,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.usa01.yuiuingkig.icu"}', 'auto'),
	(36, 'Silver-香港-BGP-03', 'zz.2hkmix.fogvip-zz.uk', '{"name":"Silver-香港-BGP-03","type":"trojan","server":"zz.2hkmix.fogvip-zz.uk","port":50306,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.2heix.yuiuingkig.icu"}', 'auto'),
	(35, 'Bronze-香港-BGP-05', 'zz.xgys.fogvip-zz.uk', '{"name":"Bronze-香港-BGP-05","type":"trojan","server":"zz.xgys.fogvip-zz.uk","port":40052,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.xgys01.yuiuingkig.icu"}', 'auto'),
	(34, 'Silver-泰国-商宽-01', 'zz.tgcloud.fogvip-zz.uk', '{"name":"Silver-泰国-商宽-01","type":"trojan","server":"zz.tgcloud.fogvip-zz.uk","port":50055,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.tgcloud.yuiuingkig.icu"}', 'auto'),
	(33, 'Silver-日本-商宽-01', 'zz.jp01.fogvip-zz.uk', '{"name":"Silver-日本-商宽-01","type":"trojan","server":"zz.jp01.fogvip-zz.uk","port":40015,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.jp01.yuiuingkig.icu"}', 'auto'),
	(32, 'Silver-马来西亚-01', 'zz.mlxy-1.fogvip-zz.uk', '{"name":"Silver-马来西亚-01","type":"trojan","server":"zz.mlxy-1.fogvip-zz.uk","port":50117,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.mlxy-cn.yuiuingkig.icu"}', 'auto'),
	(31, 'Silver-美国-BGP-02', 'zz.usa02.fogvip-zz.uk', '{"name":"Silver-美国-BGP-02","type":"trojan","server":"zz.usa02.fogvip-zz.uk","port":40045,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.meiguo02.yuiuingkig.icu"}', 'auto'),
	(30, 'Silver-美国-CN2-03', 'zz.mglsj.fogvip-zz.uk', '{"name":"Silver-美国-CN2-03","type":"trojan","server":"zz.mglsj.fogvip-zz.uk","port":50108,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.mglsj.yuiuingkig.icu"}', 'auto'),
	(29, 'Silver-韩国-商宽-01', 'zz.hg01.fogvip-zz.uk', '{"name":"Silver-韩国-商宽-01","type":"trojan","server":"zz.hg01.fogvip-zz.uk","port":40021,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.hg01.yuiuingkig.icu"}', 'auto'),
	(28, 'Bronze-香港-HGC-04', 'zz.xgb.fogvip-zz.uk', '{"name":"Bronze-香港-HGC-04","type":"trojan","server":"zz.xgb.fogvip-zz.uk","port":40025,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.xghk03.yuiuingkig.icu"}', 'auto'),
	(27, 'Silver-香港-LMT-04', 'zz.2pkshk.fogvip-zz.uk', '{"name":"Silver-香港-LMT-04","type":"trojan","server":"zz.2pkshk.fogvip-zz.uk","port":50307,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.2heix.yuiuingkig.icu"}', 'auto'),
	(26, 'Silver-新加坡-01', 'zz.xinjp02.fogvip-zz.uk', '{"name":"Silver-新加坡-01","type":"trojan","server":"zz.xinjp02.fogvip-zz.uk","port":40017,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.xinjp01.yuiuingkig.icu"}', 'auto'),
	(24, 'Silver-新加坡-02', 'zz.sghkg.fogvip-zz.uk', '{"name":"Silver-新加坡-02","type":"trojan","server":"zz.sghkg.fogvip-zz.uk","port":50311,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.xjpsghkg.yuiuingkig.icu"}', 'auto'),
	(25, 'Silver-香港-BGPZ-05', 'zz.xgbig.viptap-tcb-zz.cc', '{"name":"Silver-香港-BGPZ-05","type":"trojan","server":"zz.xgbig.viptap-tcb-zz.cc","port":50070,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.xgbig.yuiuingkig.icu"}', 'auto'),
	(45, 'Silver-台湾-商宽-01', 'zz.tw01.fogvip-zz.uk', '{"name":"Silver-台湾-商宽-01","type":"trojan","server":"zz.tw01.fogvip-zz.uk","port":40037,"password":"f43d30ee-7a5d-408f-a32e-b74935eeb7d7","sni":"ld.tw01.yuiuingkig.icu"}', 'auto');
/*!40000 ALTER TABLE `proxies` ENABLE KEYS */;

-- 导出  表 dcat-admin.sub_info 结构
CREATE TABLE IF NOT EXISTS `sub_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '名称',
  `proxies` json DEFAULT NULL COMMENT '支持的节点列表',
  `url` varchar(128) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '订阅地址',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- 正在导出表  dcat-admin.sub_info 的数据：1 rows
/*!40000 ALTER TABLE `sub_info` DISABLE KEYS */;
REPLACE INTO `sub_info` (`id`, `name`, `proxies`, `url`) VALUES
	(2, 'lieoxc', '["Bronze-香港-HGC-04", "Silver-新加坡-01", "Silver-新加坡-02"]', 'http://127.0.0.1:18080/api/conf/?token=vip1');
/*!40000 ALTER TABLE `sub_info` ENABLE KEYS */;

-- 导出  表 dcat-admin.urls 结构
CREATE TABLE IF NOT EXISTS `urls` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `urls` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '原始订阅地址',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- 正在导出表  dcat-admin.urls 的数据：1 rows
/*!40000 ALTER TABLE `urls` DISABLE KEYS */;
REPLACE INTO `urls` (`id`, `urls`) VALUES
	(1, 'https://dawson0207.xn--3iq226gfdb94q.com/api/v1/client/subscribe?token=577ea451b640cb802b44ca340a3151a6');
/*!40000 ALTER TABLE `urls` ENABLE KEYS */;

-- 导出  表 dcat-admin.users 结构
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `password` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  dcat-admin.users 的数据：0 rows
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
