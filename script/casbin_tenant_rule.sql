SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_tenant_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_tenant_rule`;
CREATE TABLE `casbin_tenant_rule`  (
  `p_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `v0` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `v1` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `v2` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `v3` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `v4` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `v5` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  INDEX `IDX_casbin_rule_p_type`(`p_type`) USING BTREE,
  INDEX `IDX_casbin_rule_v0`(`v0`) USING BTREE,
  INDEX `IDX_casbin_rule_v1`(`v1`) USING BTREE,
  INDEX `IDX_casbin_rule_v2`(`v2`) USING BTREE,
  INDEX `IDX_casbin_rule_v3`(`v3`) USING BTREE,
  INDEX `IDX_casbin_rule_v4`(`v4`) USING BTREE,
  INDEX `IDX_casbin_rule_v5`(`v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_tenant_rule
-- ----------------------------
INSERT INTO `casbin_tenant_rule` VALUES ('p', 'admin', 'tenant1', 'data1', 'read', '', '');
INSERT INTO `casbin_tenant_rule` VALUES ('p', 'admin', 'tenant2', 'data2', 'read', '', '');
INSERT INTO `casbin_tenant_rule` VALUES ('g', 'dajun', 'admin', 'tenant1', '', '', '');
INSERT INTO `casbin_tenant_rule` VALUES ('g', 'dajun', 'developer', 'tenant2', '', '', '');
INSERT INTO `casbin_tenant_rule` VALUES ('p', 'wuxiangege', 'tenant1', 'data1', 'write', '', '');

SET FOREIGN_KEY_CHECKS = 1;
