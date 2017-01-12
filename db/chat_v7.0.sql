/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50611
Source Host           : localhost:3306
Source Database       : chat

Target Server Type    : MYSQL
Target Server Version : 50611
File Encoding         : 65001

Date: 2017-01-12 18:11:42
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `chat_app_access`
-- ----------------------------
DROP TABLE IF EXISTS `chat_app_access`;
CREATE TABLE `chat_app_access` (
  `roleid` int(10) NOT NULL COMMENT '角色id',
  `nodeid` int(10) NOT NULL COMMENT '节点id',
  `level` tinyint(1) NOT NULL COMMENT '层级',
  `pid` int(10) NOT NULL COMMENT '上级',
  `module` varchar(50) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='权限表';

-- ----------------------------
-- Records of chat_app_access
-- ----------------------------
INSERT INTO `chat_app_access` VALUES ('1', '49', '3', '0', 'test');
INSERT INTO `chat_app_access` VALUES ('1', '30', '2', '0', null);
INSERT INTO `chat_app_access` VALUES ('1', '1', '1', '0', null);
INSERT INTO `chat_app_access` VALUES ('1', '39', '3', '0', null);
INSERT INTO `chat_app_access` VALUES ('1', '2', '2', '0', null);
INSERT INTO `chat_app_access` VALUES ('1', '6', '2', '0', null);
INSERT INTO `chat_app_access` VALUES ('1', '40', '2', '0', null);
INSERT INTO `chat_app_access` VALUES ('1', '50', '3', '0', null);

-- ----------------------------
-- Table structure for `chat_app_group`
-- ----------------------------
DROP TABLE IF EXISTS `chat_app_group`;
CREATE TABLE `chat_app_group` (
  `id` smallint(3) NOT NULL,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '模块名称',
  `title` varchar(30) NOT NULL COMMENT '对应解析',
  `addtime` int(11) NOT NULL,
  `status` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='应用分组';

-- ----------------------------
-- Records of chat_app_group
-- ----------------------------
INSERT INTO `chat_app_group` VALUES ('2', 'App', '应用中心', '1222841259', '1');

-- ----------------------------
-- Table structure for `chat_app_node`
-- ----------------------------
DROP TABLE IF EXISTS `chat_app_node`;
CREATE TABLE `chat_app_node` (
  `id` int(11) NOT NULL,
  `name` varchar(20) NOT NULL COMMENT '节点名称　英文',
  `title` varchar(50) NOT NULL COMMENT ' 对应中文描述',
  `status` tinyint(1) NOT NULL,
  `remark` varchar(200) NOT NULL COMMENT '描述',
  `level` tinyint(2) NOT NULL COMMENT '层级',
  `groupid` int(6) NOT NULL,
  `pid` int(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of chat_app_node
-- ----------------------------
INSERT INTO `chat_app_node` VALUES ('49', 'Read', '查看', '1', '', '3', '0', '30');
INSERT INTO `chat_app_node` VALUES ('30', 'Public', '公共模块', '1', '', '2', '0', '1');
INSERT INTO `chat_app_node` VALUES ('1', 'App', 'RBAC', '1', '', '1', '0', '0');
INSERT INTO `chat_app_node` VALUES ('39', 'List', '列表', '1', '', '3', '0', '30');
INSERT INTO `chat_app_node` VALUES ('2', 'Node', '节点管理', '1', '', '2', '0', '1');
INSERT INTO `chat_app_node` VALUES ('6', 'Role', '角色管理', '1', '', '2', '0', '1');
INSERT INTO `chat_app_node` VALUES ('40', 'Index', '默认模块', '1', '', '2', '0', '1');
INSERT INTO `chat_app_node` VALUES ('50', 'Main', '空白首页', '1', '', '3', '0', '40');

-- ----------------------------
-- Table structure for `chat_position_log`
-- ----------------------------
DROP TABLE IF EXISTS `chat_position_log`;
CREATE TABLE `chat_position_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `attime` int(11) NOT NULL,
  `nation` varchar(40) DEFAULT NULL,
  `province` varchar(20) DEFAULT NULL,
  `city` varchar(20) DEFAULT NULL,
  `adcode` int(6) DEFAULT NULL COMMENT '//行政区ID，六位数字, 前两位是省，中间是市，后面两位是区，比如深圳市ID为440300',
  `addr` varchar(100) DEFAULT NULL,
  `lat` varchar(20) DEFAULT NULL COMMENT '//火星坐标(gcj02)，腾讯、Google、高德通用',
  `lng` varchar(20) DEFAULT NULL COMMENT '//火星坐标(gcj02)，腾讯、Google、高德通用',
  `accuracy` int(8) DEFAULT NULL COMMENT '误差范围，以米为单位',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of chat_position_log
-- ----------------------------

-- ----------------------------
-- Table structure for `chat_role`
-- ----------------------------
DROP TABLE IF EXISTS `chat_role`;
CREATE TABLE `chat_role` (
  `id` int(6) NOT NULL,
  `name` varchar(20) NOT NULL,
  `pid` int(6) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `ename` varchar(5) DEFAULT NULL,
  `createtime` int(11) NOT NULL,
  `updatetime` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of chat_role
-- ----------------------------
INSERT INTO `chat_role` VALUES ('1', '领导组', '0', '1', '', '', '1484034682', '0');

-- ----------------------------
-- Table structure for `chat_role_user`
-- ----------------------------
DROP TABLE IF EXISTS `chat_role_user`;
CREATE TABLE `chat_role_user` (
  `roleid` int(11) NOT NULL,
  `userid` int(11) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of chat_role_user
-- ----------------------------
INSERT INTO `chat_role_user` VALUES ('1', '1');

-- ----------------------------
-- Table structure for `chat_user`
-- ----------------------------
DROP TABLE IF EXISTS `chat_user`;
CREATE TABLE `chat_user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` char(20) NOT NULL,
  `password` varchar(40) NOT NULL,
  `regtime` int(10) NOT NULL,
  `status` tinyint(2) NOT NULL DEFAULT '1',
  `groupid` int(5) NOT NULL DEFAULT '0',
  `email` varchar(60) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of chat_user
-- ----------------------------
INSERT INTO `chat_user` VALUES ('1', 'ming', '123456', '0', '1', '0', 'm@huiz.cn');
INSERT INTO `chat_user` VALUES ('2', 'wangming', '123456', '0', '1', '0', '');
