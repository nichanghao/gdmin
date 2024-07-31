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

 Date: 31/07/2024 13:19:37
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
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (1, 'g', 'u:1', 'r:1', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (2, 'p', 'r:1', '', '1', '', '', '');
INSERT INTO `casbin_rule` VALUES (7, 'p', 'r:1', '', '18', '', '', '');
INSERT INTO `casbin_rule` VALUES (3, 'p', 'r:1', '', '2', '', '', '');
INSERT INTO `casbin_rule` VALUES (6, 'p', 'r:1', 'sys:menu', '14', '', '', '');
INSERT INTO `casbin_rule` VALUES (8, 'p', 'r:1', 'sys:menu:add', '15', '', '', '');
INSERT INTO `casbin_rule` VALUES (9, 'p', 'r:1', 'sys:menu:delete', '17', '', '', '');
INSERT INTO `casbin_rule` VALUES (10, 'p', 'r:1', 'sys:menu:edit', '16', '', '', '');
INSERT INTO `casbin_rule` VALUES (5, 'p', 'r:1', 'sys:role', '9', '', '', '');
INSERT INTO `casbin_rule` VALUES (11, 'p', 'r:1', 'sys:role:add', '10', '', '', '');
INSERT INTO `casbin_rule` VALUES (12, 'p', 'r:1', 'sys:role:assignMenus', '12', '', '', '');
INSERT INTO `casbin_rule` VALUES (13, 'p', 'r:1', 'sys:role:delete', '13', '', '', '');
INSERT INTO `casbin_rule` VALUES (14, 'p', 'r:1', 'sys:role:edit', '11', '', '', '');
INSERT INTO `casbin_rule` VALUES (4, 'p', 'r:1', 'sys:user', '3', '', '', '');
INSERT INTO `casbin_rule` VALUES (15, 'p', 'r:1', 'sys:user:add', '4', '', '', '');
INSERT INTO `casbin_rule` VALUES (16, 'p', 'r:1', 'sys:user:assignRoles', '6', '', '', '');
INSERT INTO `casbin_rule` VALUES (17, 'p', 'r:1', 'sys:user:delete', '8', '', '', '');
INSERT INTO `casbin_rule` VALUES (19, 'p', 'r:1', 'sys:user:edit', '5', '', '', '');
INSERT INTO `casbin_rule` VALUES (18, 'p', 'r:1', 'sys:user:resetPwd', '7', '', '', '');

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '菜单名称',
  `route_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '路由名称',
  `type` tinyint(1) NULL DEFAULT NULL COMMENT '菜单类型(1:目录,2:菜单,3:按钮)',
  `permission` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '权限标识',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '路由地址',
  `component` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '组件',
  `parent_id` bigint UNSIGNED NULL DEFAULT 0 COMMENT '父菜单ID',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '菜单状态(0:禁用,1:启用)',
  `meta` json NULL COMMENT '路由元数据',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `modify_user` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '修改人',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_menu_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, '首页', 'home', 1, '', '/home', 'layout.base$view.home', 0, 1, '{\"icon\": \"mdi:monitor-dashboard\", \"order\": 1, \"i18nKey\": \"route.home\"}', '2024-07-25 10:47:55.157', '1', NULL);
INSERT INTO `sys_menu` VALUES (2, '系统管理', 'system', 1, '', '/system', 'layout.base', 0, 1, '{\"icon\": \"carbon:cloud-service-management\", \"order\": 2, \"i18nKey\": \"route.system\"}', '2024-07-25 10:50:04.754', '1', NULL);
INSERT INTO `sys_menu` VALUES (3, '用户管理', 'system_user', 2, 'sys:user', '/system/user', 'view.system_user', 2, 1, '{\"icon\": \"ic:round-manage-accounts\", \"order\": 1, \"i18nKey\": \"route.system_user\"}', '2024-07-31 13:06:24.526', '1', NULL);
INSERT INTO `sys_menu` VALUES (4, '新增用户', '', 3, 'sys:user:add', '', '', 3, 1, '{\"order\": 0, \"i18nKey\": null}', '2024-07-30 16:31:15.887', '1', NULL);
INSERT INTO `sys_menu` VALUES (5, '编辑用户', '', 3, 'sys:user:edit', '', '', 3, 1, '{\"order\": 1, \"i18nKey\": null}', '2024-07-30 16:58:40.209', '1', NULL);
INSERT INTO `sys_menu` VALUES (6, '分配角色', '', 3, 'sys:user:assignRoles', '', '', 3, 1, '{\"order\": 2, \"i18nKey\": null}', '2024-07-30 17:02:26.964', '1', NULL);
INSERT INTO `sys_menu` VALUES (7, '重置密码', '', 3, 'sys:user:resetPwd', '', '', 3, 1, '{\"order\": 4, \"i18nKey\": null}', '2024-07-30 17:04:57.943', '1', NULL);
INSERT INTO `sys_menu` VALUES (8, '删除用户', '', 3, 'sys:user:delete', '', '', 3, 1, '{\"order\": 5, \"i18nKey\": null}', '2024-07-30 17:05:44.823', '1', NULL);
INSERT INTO `sys_menu` VALUES (9, '角色管理', 'system_role', 2, 'sys:role', '/system/role', 'view.system_role', 2, 1, '{\"icon\": \"carbon:user-role\", \"order\": 2, \"i18nKey\": \"route.system_role\"}', '2024-07-31 13:12:02.585', '1', NULL);
INSERT INTO `sys_menu` VALUES (10, '新增角色', '', 3, 'sys:role:add', '', '', 9, 1, '{\"order\": 0, \"i18nKey\": null}', '2024-07-30 17:14:30.884', '1', NULL);
INSERT INTO `sys_menu` VALUES (11, '编辑角色', '', 3, 'sys:role:edit', '', '', 9, 1, '{\"order\": 1, \"i18nKey\": null}', '2024-07-30 17:15:13.891', '1', NULL);
INSERT INTO `sys_menu` VALUES (12, '分配权限', '', 3, 'sys:role:assignMenus', '', '', 9, 1, '{\"order\": 0, \"i18nKey\": null}', '2024-07-30 17:17:27.685', '1', NULL);
INSERT INTO `sys_menu` VALUES (13, '删除角色', '', 3, 'sys:role:delete', '', '', 9, 1, '{\"order\": 0, \"i18nKey\": null}', '2024-07-30 17:18:07.425', '1', NULL);
INSERT INTO `sys_menu` VALUES (14, '菜单管理', 'system_menu', 2, 'sys:menu', '/system/menu', 'view.system_menu', 2, 1, '{\"icon\": \"material-symbols:route\", \"order\": 3, \"i18nKey\": \"route.system_menu\"}', '2024-07-31 13:04:19.701', '1', NULL);
INSERT INTO `sys_menu` VALUES (15, '新增菜单', '', 3, 'sys:menu:add', '', '', 14, 1, '{\"order\": 0, \"i18nKey\": null}', '2024-07-30 17:19:10.776', '1', NULL);
INSERT INTO `sys_menu` VALUES (16, '编辑菜单', '', 3, 'sys:menu:edit', '', '', 14, 1, '{\"order\": 0, \"i18nKey\": null}', '2024-07-30 17:19:35.597', '1', NULL);
INSERT INTO `sys_menu` VALUES (17, '删除菜单', '', 3, 'sys:menu:delete', '', '', 14, 1, '{\"order\": 0, \"i18nKey\": null}', '2024-07-30 17:19:58.259', '1', NULL);
INSERT INTO `sys_menu` VALUES (18, '关于', 'about', 1, '', '/about', 'layout.base$view.about', 0, 1, '{\"icon\": \"fluent:book-information-24-regular\", \"order\": 10, \"i18nKey\": \"route.about\"}', '2024-07-31 11:18:37.165', '1', NULL);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '角色名',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '角色标识',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '状态(1:启用 2:禁用)',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `modify_user` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '修改人',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_role_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '超级管理员', 'super_admin', 1, '超级管理员', '2024-07-29 15:56:36.859', NULL, NULL);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户登录密码',
  `nickname` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户昵称',
  `gender` tinyint(1) NULL DEFAULT NULL COMMENT '性别(1:男,2:女)',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '联系电话',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '用户状态(1:正常,2:停用)',
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
INSERT INTO `sys_user` VALUES (1, 'gdmin', '$2a$10$cl.N0OlfZQPGARJrxDJpzuJ1ZnEXCAotI1o8X6yWvYC5fZihKd8Oe', 'gdmin', 1, '13800138000', 'admin@localhost', 1, '2024-07-30 17:23:59.789', NULL, NULL);

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
INSERT INTO `sys_user_role` VALUES (1, 1);

SET FOREIGN_KEY_CHECKS = 1;
