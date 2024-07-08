/*
 Navicat Premium Data Transfer

 Source Server         : mysql@localhost
 Source Server Type    : MySQL
 Source Server Version : 80200 (8.2.0)
 Source Host           : localhost:3306
 Source Schema         : gdmin

 Target Server Type    : MySQL
 Target Server Version : 80200 (8.2.0)
 File Encoding         : 65001

 Date: 08/07/2024 15:46:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
                             `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                             `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户登录名',
                             `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户登录密码',
                             `nickname` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户昵称',
                             `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '联系电话',
                             `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '邮箱',
                             `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
                             `modify_user` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '修改人',
                             `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
                             PRIMARY KEY (`id`) USING BTREE,
                             INDEX `idx_sys_user_username`(`username` ASC) USING BTREE,
                             INDEX `idx_sys_user_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '$2a$10$cl.N0OlfZQPGARJrxDJpzuJ1ZnEXCAotI1o8X6yWvYC5fZihKd8Oe', 'admin', NULL, NULL, NULL, NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;