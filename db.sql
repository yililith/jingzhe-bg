
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
                          `uid` bigint NOT NULL AUTO_INCREMENT COMMENT '用户唯一标识，自增ID从100000001开始',
                          `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名，唯一字段',
                          `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码(存储加密后的值)',
                          `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户昵称',
                          `grade` tinyint NOT NULL DEFAULT 2 COMMENT '用户等级 (1: 管理员, 2: 普通用户, 3: VIP用户等)',
                          `status` tinyint NOT NULL DEFAULT 1 COMMENT '账户状态 (0: 禁用, 1: 启用, 2: 待激活)',
                          `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                          PRIMARY KEY (`uid`) USING BTREE,
                          UNIQUE INDEX `idx_username`(`username` ASC) USING BTREE,
                          INDEX `idx_status`(`status` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 100000002 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (100000001, 'admin', '$2a$04$2jxFIjFhqXXSj5Td8DPBDOV19HyLf65OMQfVt0rwfzrfWoVRVCMnO', '超级管理员', 1, 1, '2025-07-23 17:02:30', '2025-07-23 17:02:54');

SET FOREIGN_KEY_CHECKS = 1;




/*
 Navicat Premium Dump SQL

 Date: 23/07/2025 17:36:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_menu
-- ----------------------------
DROP TABLE IF EXISTS `user_menu`;
CREATE TABLE `user_menu`  (
                              `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                              `mid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单ID',
                              `uid` char(9) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户ID，9位数字',
                              `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否启用：1启用，0禁用',
                              `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                              PRIMARY KEY (`id`) USING BTREE,
                              UNIQUE INDEX `uk_mid_uid`(`mid` ASC, `uid` ASC) USING BTREE,
                              INDEX `idx_uid`(`uid` ASC) USING BTREE,
                              INDEX `idx_status`(`status` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户菜单关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_menu
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;


/*
 Navicat Premium Dump SQL


 Date: 23/07/2025 17:37:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_image
-- ----------------------------
DROP TABLE IF EXISTS `user_image`;
CREATE TABLE `user_image`  (
                               `IID` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '图片唯一ID',
                               `UID` char(9) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户ID，9位数字',
                               `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '图片在OSS中的路径',
                               `bucket` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '图片所在存储桶',
                               `is_avatar` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否为头像：1是，0不是',
                               `file_size` int NOT NULL COMMENT '图片大小(字节)',
                               `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否逻辑删除：1是，0不是',
                               `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                               PRIMARY KEY (`IID`) USING BTREE,
                               INDEX `idx_uid`(`UID` ASC) USING BTREE,
                               INDEX `idx_is_avatar`(`is_avatar` ASC) USING BTREE,
                               INDEX `idx_is_deleted`(`is_deleted` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户图片关联表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;



/*
 Navicat Premium Dump SQL


 Date: 23/07/2025 17:38:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus`  (
                          `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单唯一ID',
                          `menu_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单名称',
                          `menu_name_cn` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单中文名称',
                          `menu_level` tinyint(1) NOT NULL COMMENT '菜单等级：1为主菜单，2为子菜单',
                          `menu_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单类型',
                          `menu_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '菜单状态：1启用，0禁用',
                          `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                          PRIMARY KEY (`id`) USING BTREE,
                          INDEX `idx_menu_level`(`menu_level` ASC) USING BTREE,
                          INDEX `idx_menu_status`(`menu_status` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '系统菜单表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
