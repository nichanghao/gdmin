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

 Date: 26/07/2024 17:40:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype` ASC, `v0` ASC, `v1` ASC, `v2` ASC, `v3` ASC, `v4` ASC, `v5` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 51 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (45, 'g', 'u:1', 'r:2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (46, 'p', 'r:2', '', '4', '', '', '');
INSERT INTO `casbin_rule` VALUES (47, 'p', 'r:2', '', '5', '', '', '');
INSERT INTO `casbin_rule` VALUES (48, 'p', 'r:2', '', '6', '', '', '');
INSERT INTO `casbin_rule` VALUES (49, 'p', 'r:2', '', '7', '', '', '');
INSERT INTO `casbin_rule` VALUES (50, 'p', 'r:2', '', '8', '', '', '');

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '菜单名称',
  `route_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `type` tinyint(1) NULL DEFAULT NULL COMMENT '菜单类型(1:目录,2:菜单)',
  `permission` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '权限标识',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '路由地址',
  `component` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '组件',
  `parent_id` bigint UNSIGNED NULL DEFAULT 0 COMMENT '父菜单ID',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '菜单状态(0:禁用,1:启用)',
  `meta` json NULL COMMENT '菜单元数据',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `modify_user` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '修改人',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_menu_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (3, '首页', 'home', 1, '', '/home', 'layout.base$view.home', 0, 1, '{\"icon\": \"mdi:monitor-dashboard\", \"order\": 2, \"i18nKey\": \"route.home\"}', '2024-07-24 16:59:37.945', '1', '2024-07-24 17:21:23.898');
INSERT INTO `sys_menu` VALUES (4, '首页', 'home', 1, '', '/home', 'layout.base$view.home', 0, 1, '{\"icon\": \"mdi:monitor-dashboard\", \"order\": 1, \"i18nKey\": \"route.home\"}', '2024-07-25 10:47:55.157', '1', NULL);
INSERT INTO `sys_menu` VALUES (5, '系统管理', 'system', 1, '', '/system', 'layout.base', 0, 1, '{\"icon\": \"carbon:cloud-service-management\", \"order\": 2, \"i18nKey\": \"route.manage\"}', '2024-07-25 10:50:04.754', '1', NULL);
INSERT INTO `sys_menu` VALUES (6, '用户管理', 'system_user', 2, '', '/system/user', 'view.system_user', 5, 1, '{\"icon\": \"ic:round-manage-accounts\", \"order\": 1, \"i18nKey\": \"route.manage_user\"}', '2024-07-25 10:52:10.855', '1', NULL);
INSERT INTO `sys_menu` VALUES (7, '角色管理', 'system_role', 2, '', '/system/role', 'view.system_role', 5, 1, '{\"icon\": \"carbon:user-role\", \"order\": 2, \"i18nKey\": \"route.manage_role\"}', '2024-07-25 10:54:00.997', '1', NULL);
INSERT INTO `sys_menu` VALUES (8, '菜单管理', 'system_menu', 2, '', '/system/menu', 'view.system_menu', 5, 1, '{\"icon\": \"material-symbols:route\", \"order\": 3, \"i18nKey\": \"route.manage_menu\"}', '2024-07-25 10:55:10.461', '1', NULL);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '角色名',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '角色标识',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `modify_user` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '修改人',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_role_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '管理员', 'super_admin11', '系统管理员', '2024-07-05 11:23:31.778', '', '2024-07-05 11:23:55.058');
INSERT INTO `sys_role` VALUES (2, '超级管理员', 'super_admin', '超级管理员', '2024-07-05 11:04:19.112', '', NULL);
INSERT INTO `sys_role` VALUES (13, '管理员1', 'admin1', '管理员1', '2024-07-19 16:34:52.301', '1', NULL);

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
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '用户状态(1:正常,2:停用)',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_user_username`(`username` ASC) USING BTREE,
  INDEX `idx_sys_user_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'gdmin', '$2a$10$cl.N0OlfZQPGARJrxDJpzuJ1ZnEXCAotI1o8X6yWvYC5fZihKd8Oe', 'admin', '13800138000', 'admin@localhost', '2024-07-19 10:46:16.305', NULL, NULL, 1);
INSERT INTO `sys_user` VALUES (2, 'admin22', '$2a$10$cl.N0OlfZQPGARJrxDJpzuJ1ZnEXCAotI1o8X6yWvYC5fZihKd8Oe', 'admin22', '13800138000', 'admin@localhost', '2024-07-10 18:00:13.938', NULL, NULL, 1);

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `sys_role_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  `sys_user_id` bigint UNSIGNED NOT NULL COMMENT '用户ID',
  PRIMARY KEY (`sys_role_id`, `sys_user_id`) USING BTREE,
  INDEX `fk_sys_user_role_sys_user`(`sys_user_id` ASC) USING BTREE,
  CONSTRAINT `fk_sys_user_role_sys_role` FOREIGN KEY (`sys_role_id`) REFERENCES `sys_role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_user_role_sys_user` FOREIGN KEY (`sys_user_id`) REFERENCES `sys_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (2, 1);
INSERT INTO `sys_user_role` VALUES (1, 2);

SET FOREIGN_KEY_CHECKS = 1;
