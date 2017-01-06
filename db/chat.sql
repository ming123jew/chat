/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50611
Source Host           : localhost:3306
Source Database       : chat

Target Server Type    : MYSQL
Target Server Version : 50611
File Encoding         : 65001

Date: 2017-01-06 17:38:49
*/

SET FOREIGN_KEY_CHECKS=0;

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
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of chat_user
-- ----------------------------
INSERT INTO `chat_user` VALUES ('1', 'ming', '123456', '0', '1', '0', 'm@huiz.cn');
