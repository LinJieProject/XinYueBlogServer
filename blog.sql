/*
 Navicat Premium Data Transfer

 Source Server         : MySQL
 Source Server Type    : MySQL
 Source Server Version : 80018
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80018
 File Encoding         : 65001

 Date: 24/10/2019 10:52:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `tag_id` bigint(20) unsigned NOT NULL COMMENT '标签ID',
  `content` longtext COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章内容',
  `title` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标题',
  `view_conut` int(255) unsigned NOT NULL COMMENT '阅读次数',
  `comment_conut` int(255) unsigned NOT NULL COMMENT '评论次数',
  `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章作者',
  `summary` varchar(256) COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章摘要',
  `create_time` timestamp NOT NULL COMMENT '发布时间',
  `update_time` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '表ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `gender` tinyint(4) NOT NULL DEFAULT '0' COMMENT '性别',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
