/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : 127.0.0.1:3306
 Source Schema         : ginvel_test

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 19/07/2022 18:08:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gv_admin
-- ----------------------------
DROP TABLE IF EXISTS `gv_admin`;
CREATE TABLE `gv_admin` (
  `admin_id` int NOT NULL AUTO_INCREMENT,
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `avatar_img` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `update_time` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `create_time` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `update_ip` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ipv4 、ipv6',
  `state` int NOT NULL DEFAULT '1' COMMENT '1可用，2删除',
  `login_level` int NOT NULL DEFAULT '10' COMMENT '10超级管理员，12组长，14普通成员，16游客',
  `login_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `login_pwd` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `login_phone` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `login_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`admin_id`),
  UNIQUE KEY `admin_id` (`admin_id`,`update_ip`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of gv_admin
-- ----------------------------
BEGIN;
INSERT INTO `gv_admin` VALUES (1, 'superuser', 'https://hao123.s3.cn-north-1.jdcloud-oss.com/static/img/zhuanzhuanle/not.png', '20220531150101', '20220601145107', '127.0.0.1', 1, 10, 'superfYone2022', 'c96eda712faffaf4d19aba33de6ffa09', NULL, NULL);
INSERT INTO `gv_admin` VALUES (2, 'supertest', NULL, '20220601162507', '20220601145107', '127.0.0.1', 1, 10, 'superTestfY2022', 'f6c84d742f04e47f75bf6252f65ca254', NULL, NULL);
INSERT INTO `gv_admin` VALUES (3, 'postman01', NULL, NULL, '20220601150453', '127.0.0.1', 1, 10, 'postmanTest01', 'c8048b7820c927b54e4be0891949aad9', NULL, NULL);
INSERT INTO `gv_admin` VALUES (4, 'postman02', NULL, NULL, '20220601151309', '127.0.0.1', 2, 10, 'postmanteSt02', 'ecad8fd043de9e79d623d3fbf46e6143', NULL, NULL);
INSERT INTO `gv_admin` VALUES (5, NULL, NULL, NULL, '20220601153159', '127.0.0.1', 2, 10, 'gotest02', '486a34d1f1d205ea85dd063a4013b6ce', NULL, NULL);
INSERT INTO `gv_admin` VALUES (6, NULL, NULL, NULL, '20220601153207', '127.0.0.1', 2, 10, 'gotest02', '486a34d1f1d205ea85dd063a4013b6ce', NULL, NULL);
INSERT INTO `gv_admin` VALUES (7, NULL, NULL, NULL, '20220601153308', '127.0.0.1', 2, 10, 'gotest02', '486a34d1f1d205ea85dd063a4013b6ce', NULL, NULL);
INSERT INTO `gv_admin` VALUES (8, NULL, NULL, NULL, '20220601153616', '127.0.0.1', 2, 10, 'gotest02', '486a34d1f1d205ea85dd063a4013b6ce', NULL, NULL);
INSERT INTO `gv_admin` VALUES (9, NULL, NULL, NULL, '20220601153747', '127.0.0.1', 2, 10, 'gotest02', '486a34d1f1d205ea85dd063a4013b6ce', NULL, NULL);
INSERT INTO `gv_admin` VALUES (10, NULL, NULL, NULL, '20220601153824', '127.0.0.1', 2, 10, 'gotest02', '486a34d1f1d205ea85dd063a4013b6ce', NULL, NULL);
INSERT INTO `gv_admin` VALUES (11, NULL, NULL, NULL, '20220601153824', '127.0.0.1', 1, 10, 'gotest02', '486a34d1f1d205ea85dd063a4013b6ce', NULL, NULL);
INSERT INTO `gv_admin` VALUES (12, NULL, NULL, NULL, '20220601153845', '127.0.0.1', 1, 10, 'gotest02', '486a34d1f1d205ea85dd063a4013b6ce', NULL, NULL);
INSERT INTO `gv_admin` VALUES (13, NULL, NULL, '20220601155754', '20220601154327', '127.0.0.1', 1, 10, 'gotest031', '#other#', NULL, NULL);
INSERT INTO `gv_admin` VALUES (14, NULL, NULL, '20220601161834', '20220601155115', '127.0.0.1', 1, 10, 'gotest01', '123456', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for gv_admin_token
-- ----------------------------
DROP TABLE IF EXISTS `gv_admin_token`;
CREATE TABLE `gv_admin_token` (
  `admin_token_id` int NOT NULL AUTO_INCREMENT,
  `admin_id` int NOT NULL,
  `login_token` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `update_time` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `update_ip` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'IPv4或IPv6',
  `app_class` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录客户端的种类',
  `ua` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ua',
  `screen_info` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '屏幕像素',
  `refer` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录页来源',
  PRIMARY KEY (`admin_token_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of gv_admin_token
-- ----------------------------
BEGIN;
INSERT INTO `gv_admin_token` VALUES (16, 2, 'YuiLpiuHktvo53eG55r7bojttZpAXJu4Hjokp2VIwosc2nMzKu/k5N5Wopip+AEu', '20220624183616', '127.0.0.1', 'test', 'null key', NULL, NULL);
INSERT INTO `gv_admin_token` VALUES (18, 2, 'YuiLpiuHktvhwKK7lFA7/li+HKgvVKBvxLH4asPMYoOnjrFRxuH0emYWzfn39oAI', '20220624183622', '127.0.0.1', 'test', 'null key', NULL, NULL);
INSERT INTO `gv_admin_token` VALUES (19, 2, 'YuiLpiuHkttf+OLWR8HDnjeLJFyshQTCH0lWhQtS2dqVjrxYDKlAqXv0+Rj4P5bI', '20220624182517', '127.0.0.1', 'test', NULL, NULL, NULL);
INSERT INTO `gv_admin_token` VALUES (20, 2, 'YuiLpiuHktvwz+XP4OZ+V/p9a/BhFDafHOaR8tdvX0Gz9PeqzUV2oxgXIX90C4+j', '20220624182523', '127.0.0.1', 'test', NULL, NULL, NULL);
INSERT INTO `gv_admin_token` VALUES (21, 2, 'YuiLpiuHktsjtyFuL5n85QRKEz7M5PuDl7tJYMUEGH7iHfWIo128QIAEDIDc8VE6', '20220624182529', '127.0.0.1', 'test', NULL, NULL, NULL);
INSERT INTO `gv_admin_token` VALUES (22, 2, 'YuiLpiuHktu2gYUJF+n+AdEB6iRdLdncNlrlcw2K/WNE53LXvY0W2Vk56KjoolE2', '20220624182530', '127.0.0.1', 'test', NULL, NULL, NULL);
INSERT INTO `gv_admin_token` VALUES (23, 2, 'YuiLpiuHktvML6RmOCPLT7Gz9ImHqZ/2I/l134g5f7vCa+t8LYJF/KY3BCXufr64', '20220624182538', '127.0.0.1', 'test', NULL, NULL, NULL);
INSERT INTO `gv_admin_token` VALUES (24, 3, 'YuiLpiuHktvHsLG3A2fPcQskohe7aXlF0uLzNSYBVJij+A1l1Z23sI6JHLmfeX2u', '20220624184138', '127.0.0.1', 'test', NULL, NULL, NULL);
INSERT INTO `gv_admin_token` VALUES (25, 1, '7VTCsAy+cfJRYQWov0IDAC/LgQPkgVdCsnNRNr9+xr3qBR5lrpF2nHIrZ3RuYKmI', '20220625111403', '127.0.0.1', 'vue-example', '<nil>', NULL, NULL);
INSERT INTO `gv_admin_token` VALUES (26, 1, 'Yelh88F5VeLAjmI0gigVy78qW5bz0n+dEwrBOTzVlWX7/GkQX0M2ym1r365LgJHf', '20220626222110', '127.0.0.1', 'vue-example', '<nil>', NULL, NULL);
INSERT INTO `gv_admin_token` VALUES (27, 1, 'mSx96gANcWjMtafLwzJ+JvvQLDstcy0B8N8S6VOmo+MlQ8QOzDHDjv3WUFbPSScm', '20220627013022', '127.0.0.1', 'vue-example', '<nil>', NULL, NULL);
INSERT INTO `gv_admin_token` VALUES (28, 1, '7VTCsAy+cfKCsMoV4IAfRiG4HGdo3Zetc3XjYfZTIn67X/2+wNOkFjEpqmYq7KZ9', '20220625020735', '127.0.0.1', 'vue-example', '<nil>', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for gv_captcha
-- ----------------------------
DROP TABLE IF EXISTS `gv_captcha`;
CREATE TABLE `gv_captcha` (
  `captcha_id` int NOT NULL AUTO_INCREMENT,
  `uuid` char(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `captcha` char(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `use_num` int NOT NULL DEFAULT '0' COMMENT '使用或查询过的次数',
  `create_time` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`captcha_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of gv_captcha
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
