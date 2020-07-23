/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50641
Source Host           : localhost:3306
Source Database       : yuepu

Target Server Type    : MYSQL
Target Server Version : 50641
File Encoding         : 65001

Date: 2020-07-21 18:25:34
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for area_type
-- ----------------------------
DROP TABLE IF EXISTS `area_type`;
CREATE TABLE `area_type` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `min_area` double DEFAULT NULL,
  `max_area` double DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of area_type
-- ----------------------------
INSERT INTO `area_type` VALUES ('6', '2020-07-16 16:25:24', '2020-07-16 16:34:02', null, '0', '5');
INSERT INTO `area_type` VALUES ('7', '2020-07-16 16:29:25', '2020-07-16 16:29:25', null, '1', '50');
INSERT INTO `area_type` VALUES ('8', '2020-07-16 16:33:06', '2020-07-16 16:33:06', null, '50', '100');
INSERT INTO `area_type` VALUES ('9', '2020-07-16 16:33:23', '2020-07-16 16:33:23', null, '100', '150');

-- ----------------------------
-- Table structure for city
-- ----------------------------
DROP TABLE IF EXISTS `city`;
CREATE TABLE `city` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(10) NOT NULL,
  `code` varchar(10) NOT NULL,
  `province_code` varchar(10) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of city
-- ----------------------------
INSERT INTO `city` VALUES ('1', null, null, null, '佛山市', '0001', '001');
INSERT INTO `city` VALUES ('2', null, null, null, '中山市', '0002', '001');
INSERT INTO `city` VALUES ('3', null, null, null, '广州市', '0003', '001');

-- ----------------------------
-- Table structure for district
-- ----------------------------
DROP TABLE IF EXISTS `district`;
CREATE TABLE `district` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(10) NOT NULL,
  `code` varchar(10) NOT NULL,
  `city_code` varchar(10) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of district
-- ----------------------------
INSERT INTO `district` VALUES ('1', null, null, null, '禅城区', '00001', '0001');
INSERT INTO `district` VALUES ('2', null, null, null, '南海区', '00002', '0001');
INSERT INTO `district` VALUES ('3', null, null, null, '三水区', '00003', '0001');
INSERT INTO `district` VALUES ('8', null, null, null, '顺德区', '00004', '0001');
INSERT INTO `district` VALUES ('9', null, null, null, '高明区', '00005', '0001');
INSERT INTO `district` VALUES ('10', null, null, null, '石歧', '00006', '0002');
INSERT INTO `district` VALUES ('11', null, null, null, '坦洲', '00007', '0002');
INSERT INTO `district` VALUES ('12', null, null, null, '火炬开发区', '00008', '0002');
INSERT INTO `district` VALUES ('13', null, null, null, '东区', '00009', '0002');
INSERT INTO `district` VALUES ('14', null, null, null, '西区', '00010', '0002');
INSERT INTO `district` VALUES ('15', null, null, null, '小榄', '00011', '0002');
INSERT INTO `district` VALUES ('16', null, null, null, '三乡', '00012', '0002');
INSERT INTO `district` VALUES ('17', null, null, null, '沙溪', '00013', '0002');
INSERT INTO `district` VALUES ('18', null, null, null, '古镇', '00014', '0002');
INSERT INTO `district` VALUES ('19', null, null, null, '南区', '00015', '0002');
INSERT INTO `district` VALUES ('20', null, null, null, '港口', '00016', '0002');
INSERT INTO `district` VALUES ('21', null, null, null, '东升', '00017', '0002');
INSERT INTO `district` VALUES ('22', null, null, null, '南朗', '00018', '0002');
INSERT INTO `district` VALUES ('23', null, null, null, '东凤', '00019', '0002');
INSERT INTO `district` VALUES ('24', null, null, null, '横栏', '00020', '0002');
INSERT INTO `district` VALUES ('25', null, null, null, '南头', '00021', '0002');
INSERT INTO `district` VALUES ('26', null, null, null, '黄圃', '00022', '0002');
INSERT INTO `district` VALUES ('27', null, null, null, '三角', '00023', '0002');
INSERT INTO `district` VALUES ('28', null, null, null, '板芙', '00024', '0002');
INSERT INTO `district` VALUES ('29', null, null, null, '五桂山', '00025', '0002');
INSERT INTO `district` VALUES ('30', null, null, null, '大涌', '00026', '0002');
INSERT INTO `district` VALUES ('31', null, null, null, '民众', '00027', '0002');
INSERT INTO `district` VALUES ('32', null, null, null, '阜沙', '00028', '0002');
INSERT INTO `district` VALUES ('33', null, null, null, '神湾', '00029', '0002');
INSERT INTO `district` VALUES ('34', null, null, null, '白云区', '00030', '0003');
INSERT INTO `district` VALUES ('35', null, null, null, '天河区', '00031', '0003');
INSERT INTO `district` VALUES ('36', null, null, null, '番禺区', '00032', '0003');
INSERT INTO `district` VALUES ('37', null, null, null, '海珠区', '00033', '0003');
INSERT INTO `district` VALUES ('40', null, null, null, '花都区', '00040', '0003');

-- ----------------------------
-- Table structure for industry
-- ----------------------------
DROP TABLE IF EXISTS `industry`;
CREATE TABLE `industry` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(20) NOT NULL,
  `sort` bigint(20) DEFAULT NULL,
  `is_enable` tinyint(1) DEFAULT NULL,
  `parent_id` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of industry
-- ----------------------------
INSERT INTO `industry` VALUES ('8', null, null, null, '酒店餐饮', '99', '1', '0');
INSERT INTO `industry` VALUES ('9', null, null, null, '餐馆', '20', '1', '8');
INSERT INTO `industry` VALUES ('10', null, null, null, '冷饮甜品', '19', '1', '8');
INSERT INTO `industry` VALUES ('11', null, null, null, '面包店', '18', '1', '8');
INSERT INTO `industry` VALUES ('12', null, null, null, '食堂', '17', '1', '8');
INSERT INTO `industry` VALUES ('13', null, null, null, '咖啡馆', '16', '1', '8');
INSERT INTO `industry` VALUES ('14', null, null, null, '茶艺馆', '15', '1', '8');
INSERT INTO `industry` VALUES ('15', null, null, null, '早餐小吃店', '14', '1', '8');
INSERT INTO `industry` VALUES ('16', null, null, null, '快餐店', '13', '1', '8');
INSERT INTO `industry` VALUES ('17', null, null, null, '西餐厅', '12', '1', '8');
INSERT INTO `industry` VALUES ('18', null, null, null, '火锅店', '11', '1', '8');
INSERT INTO `industry` VALUES ('19', null, null, null, '大排档', '10', '1', '8');
INSERT INTO `industry` VALUES ('20', null, null, null, '烧烤店', '9', '1', '8');
INSERT INTO `industry` VALUES ('21', null, null, null, '茶楼', '8', '1', '8');
INSERT INTO `industry` VALUES ('22', null, null, null, '面馆', '7', '1', '8');
INSERT INTO `industry` VALUES ('23', null, null, null, '火锅干锅', '6', '1', '8');
INSERT INTO `industry` VALUES ('24', null, null, null, '其他', '5', '1', '8');
INSERT INTO `industry` VALUES ('25', null, null, null, '农家乐', '4', '1', '8');
INSERT INTO `industry` VALUES ('26', null, null, null, '美容美发', '99', '1', '0');
INSERT INTO `industry` VALUES ('27', null, null, null, '母婴养生馆', '20', '1', '26');
INSERT INTO `industry` VALUES ('28', null, null, null, '养生馆', '19', '1', '26');
INSERT INTO `industry` VALUES ('29', null, null, null, '产后修复', '18', '1', '26');
INSERT INTO `industry` VALUES ('30', null, null, null, 'SPA馆', '17', '1', '26');
INSERT INTO `industry` VALUES ('31', null, null, null, '美甲店', '16', '1', '26');
INSERT INTO `industry` VALUES ('32', null, null, null, '美发店', '15', '1', '26');
INSERT INTO `industry` VALUES ('33', null, null, null, '美容院', '14', '1', '26');
INSERT INTO `industry` VALUES ('34', null, null, null, '瑜伽馆', '13', '1', '26');
INSERT INTO `industry` VALUES ('35', null, null, null, '百货超市', '99', '1', '0');
INSERT INTO `industry` VALUES ('36', null, null, null, '水果店', '20', '1', '35');
INSERT INTO `industry` VALUES ('37', null, null, null, '超市', '19', '1', '35');
INSERT INTO `industry` VALUES ('38', null, null, null, '文具店', '18', '1', '35');
INSERT INTO `industry` VALUES ('39', null, null, null, '玩具店', '17', '1', '35');
INSERT INTO `industry` VALUES ('40', null, null, null, '母婴用品店', '16', '1', '35');
INSERT INTO `industry` VALUES ('41', null, null, null, '烟酒茶叶店', '15', '1', '35');
INSERT INTO `industry` VALUES ('42', null, null, null, '杂货店', '14', '1', '35');
INSERT INTO `industry` VALUES ('43', null, null, null, '便利店', '13', '1', '35');
INSERT INTO `industry` VALUES ('44', null, null, null, '眼睛店', '12', '1', '35');
INSERT INTO `industry` VALUES ('45', null, null, null, '化妆品店', '11', '1', '35');
INSERT INTO `industry` VALUES ('46', null, null, null, '乐器店', '10', '1', '35');
INSERT INTO `industry` VALUES ('47', null, null, null, '副食品店', '9', '1', '35');
INSERT INTO `industry` VALUES ('48', null, null, null, '档口摊位', '8', '1', '35');
INSERT INTO `industry` VALUES ('49', null, null, null, '特产类', '7', '1', '35');
INSERT INTO `industry` VALUES ('50', null, null, null, '水产肉类熟食', '6', '1', '35');
INSERT INTO `industry` VALUES ('51', null, null, null, '床上用品', '5', '1', '35');
INSERT INTO `industry` VALUES ('52', null, null, null, '休闲食品', '4', '1', '35');
INSERT INTO `industry` VALUES ('53', null, null, null, '工艺品店', '3', '1', '35');
INSERT INTO `industry` VALUES ('54', null, null, null, '书店', '2', '1', '35');
INSERT INTO `industry` VALUES ('55', null, null, null, '酒庄', '1', '1', '35');
INSERT INTO `industry` VALUES ('56', null, null, null, '服饰鞋包', '99', '1', '0');
INSERT INTO `industry` VALUES ('57', null, null, null, '专柜', '20', '1', '56');
INSERT INTO `industry` VALUES ('58', null, null, null, '精品店', '19', '1', '56');
INSERT INTO `industry` VALUES ('59', null, null, null, '皮具护理', '18', '1', '56');
INSERT INTO `industry` VALUES ('60', null, null, null, '内衣店', '17', '1', '56');
INSERT INTO `industry` VALUES ('61', null, null, null, '童装店', '16', '1', '56');
INSERT INTO `industry` VALUES ('62', null, null, null, '鞋店', '15', '1', '56');
INSERT INTO `industry` VALUES ('63', null, null, null, '箱包店', '14', '1', '56');
INSERT INTO `industry` VALUES ('64', null, null, null, '饰品店', '13', '1', '56');
INSERT INTO `industry` VALUES ('65', null, null, null, '黄金珠宝', '12', '1', '56');
INSERT INTO `industry` VALUES ('66', null, null, null, '格子铺', '11', '1', '56');
INSERT INTO `industry` VALUES ('67', null, null, null, '服装店', '10', '1', '56');
INSERT INTO `industry` VALUES ('68', null, null, null, '休闲娱乐', '99', '1', '0');
INSERT INTO `industry` VALUES ('69', null, null, null, '度假山庄', '20', '1', '68');
INSERT INTO `industry` VALUES ('70', null, null, null, '夜总会', '19', '1', '68');
INSERT INTO `industry` VALUES ('71', null, null, null, '歌舞厅（KTV）', '18', '1', '68');
INSERT INTO `industry` VALUES ('72', null, null, null, '麻将馆', '17', '1', '68');
INSERT INTO `industry` VALUES ('73', null, null, null, '球馆', '16', '1', '68');
INSERT INTO `industry` VALUES ('74', null, null, null, '水疗', '15', '1', '68');
INSERT INTO `industry` VALUES ('75', null, null, null, '足浴', '14', '1', '68');
INSERT INTO `industry` VALUES ('76', null, null, null, '酒吧', '13', '1', '68');
INSERT INTO `industry` VALUES ('77', null, null, null, '桌球城', '12', '1', '68');
INSERT INTO `industry` VALUES ('78', null, null, null, '健身房', '11', '1', '68');
INSERT INTO `industry` VALUES ('79', null, null, null, '休闲中心', '10', '1', '68');
INSERT INTO `industry` VALUES ('80', null, null, null, '游乐场', '9', '1', '68');
INSERT INTO `industry` VALUES ('81', null, null, null, '古玩字画', '8', '1', '68');
INSERT INTO `industry` VALUES ('82', null, null, null, '电影院', '7', '1', '68');
INSERT INTO `industry` VALUES ('83', null, null, null, '古玩城', '6', '1', '68');
INSERT INTO `industry` VALUES ('84', null, null, null, '电玩城', '5', '1', '68');
INSERT INTO `industry` VALUES ('85', null, null, null, '溜冰场', '4', '1', '68');
INSERT INTO `industry` VALUES ('86', null, null, null, '桌游', '3', '1', '68');
INSERT INTO `industry` VALUES ('87', null, null, null, '浴场浴室', '2', '1', '68');
INSERT INTO `industry` VALUES ('88', null, null, null, '棋牌室', '1', '1', '68');
INSERT INTO `industry` VALUES ('89', null, null, null, '网吧', '1', '1', '68');

-- ----------------------------
-- Table structure for province
-- ----------------------------
DROP TABLE IF EXISTS `province`;
CREATE TABLE `province` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(10) NOT NULL,
  `code` varchar(10) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of province
-- ----------------------------
INSERT INTO `province` VALUES ('1', null, null, null, '广东省', '001');

-- ----------------------------
-- Table structure for street
-- ----------------------------
DROP TABLE IF EXISTS `street`;
CREATE TABLE `street` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(10) NOT NULL,
  `code` varchar(10) NOT NULL,
  `district_code` varchar(10) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`),
  KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=131 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of street
-- ----------------------------
INSERT INTO `street` VALUES ('1', null, null, null, '南庄', '000001', '00001');
INSERT INTO `street` VALUES ('2', null, null, null, '张槎', '000002', '00001');
INSERT INTO `street` VALUES ('5', null, null, null, '祖庙', '000003', '00001');
INSERT INTO `street` VALUES ('6', null, null, null, '石湾', '000004', '00001');
INSERT INTO `street` VALUES ('7', null, null, null, '魁奇', '000005', '00001');
INSERT INTO `street` VALUES ('8', null, null, null, '普澜', '000006', '00001');
INSERT INTO `street` VALUES ('9', null, null, null, '弼塘', '000007', '00001');
INSERT INTO `street` VALUES ('10', null, null, null, '亲仁', '000008', '00001');
INSERT INTO `street` VALUES ('11', null, null, null, '唐园', '000009', '00001');
INSERT INTO `street` VALUES ('12', null, null, null, '绿景', '000010', '00001');
INSERT INTO `street` VALUES ('13', null, null, null, '福升', '000011', '00001');
INSERT INTO `street` VALUES ('14', null, null, null, '湖景', '000012', '00001');
INSERT INTO `street` VALUES ('15', null, null, null, '季华路', '000013', '00001');
INSERT INTO `street` VALUES ('16', null, null, null, '岭南大道', '000014', '00001');
INSERT INTO `street` VALUES ('17', null, null, null, '汾江南路', '000015', '00001');
INSERT INTO `street` VALUES ('18', null, null, null, '禅城区周边', '000016', '00001');
INSERT INTO `street` VALUES ('19', null, null, null, '环市', '000017', '00001');
INSERT INTO `street` VALUES ('20', null, null, null, '普君', '000018', '00001');
INSERT INTO `street` VALUES ('21', null, null, null, '东方广场', '000019', '00001');
INSERT INTO `street` VALUES ('22', null, null, null, '朝安', '000020', '00001');
INSERT INTO `street` VALUES ('23', null, null, null, '里水', '000021', '00002');
INSERT INTO `street` VALUES ('24', null, null, null, '大沥', '000022', '00002');
INSERT INTO `street` VALUES ('25', null, null, null, '丹灶', '000023', '00002');
INSERT INTO `street` VALUES ('26', null, null, null, '西樵', '000024', '00002');
INSERT INTO `street` VALUES ('27', null, null, null, '九江镇', '000025', '00002');
INSERT INTO `street` VALUES ('28', null, null, null, '罗村', '000026', '00002');
INSERT INTO `street` VALUES ('29', null, null, null, '松岗', '000027', '00002');
INSERT INTO `street` VALUES ('30', null, null, null, '狮山', '000028', '00002');
INSERT INTO `street` VALUES ('31', null, null, null, '黄岐', '000029', '00002');
INSERT INTO `street` VALUES ('32', null, null, null, '千灯湖', '000030', '00002');
INSERT INTO `street` VALUES ('33', null, null, null, '平洲', '000031', '00002');
INSERT INTO `street` VALUES ('34', null, null, null, '盐步', '000032', '00002');
INSERT INTO `street` VALUES ('35', null, null, null, '南海广场', '000033', '00002');
INSERT INTO `street` VALUES ('36', null, null, null, '天佑', '000034', '00002');
INSERT INTO `street` VALUES ('39', null, null, null, '朝安', '000035', '00002');
INSERT INTO `street` VALUES ('40', null, null, null, '南海大道北', '000036', '00002');
INSERT INTO `street` VALUES ('41', null, null, null, '城市广场', '000037', '00002');
INSERT INTO `street` VALUES ('42', null, null, null, '桂城', '000038', '00002');
INSERT INTO `street` VALUES ('43', null, null, null, '顺德碧桂园', '000039', '00004');
INSERT INTO `street` VALUES ('44', null, null, null, '均安', '000040', '00004');
INSERT INTO `street` VALUES ('45', null, null, null, '杏坛', '000041', '00004');
INSERT INTO `street` VALUES ('46', null, null, null, '龙江', '000042', '00004');
INSERT INTO `street` VALUES ('47', null, null, null, '乐从', '000043', '00004');
INSERT INTO `street` VALUES ('48', null, null, null, '北滘', '000044', '00004');
INSERT INTO `street` VALUES ('49', null, null, null, '陈村', '000045', '00004');
INSERT INTO `street` VALUES ('50', null, null, null, '勒流', '000046', '00004');
INSERT INTO `street` VALUES ('51', null, null, null, '伦教', '000047', '00004');
INSERT INTO `street` VALUES ('52', null, null, null, '容桂', '000048', '00004');
INSERT INTO `street` VALUES ('53', null, null, null, '大良', '000049', '00004');
INSERT INTO `street` VALUES ('54', null, null, null, '南山镇', '000050', '00003');
INSERT INTO `street` VALUES ('55', null, null, null, '芦苞', '000051', '00003');
INSERT INTO `street` VALUES ('56', null, null, null, '白坭', '000052', '00003');
INSERT INTO `street` VALUES ('57', null, null, null, '乐平', '000053', '00003');
INSERT INTO `street` VALUES ('58', null, null, null, '大塘', '000054', '00003');
INSERT INTO `street` VALUES ('59', null, null, null, '西南', '000055', '00003');
INSERT INTO `street` VALUES ('60', null, null, null, '更合', '000056', '00005');
INSERT INTO `street` VALUES ('61', null, null, null, '富湾', '000057', '00005');
INSERT INTO `street` VALUES ('62', null, null, null, '明城', '000058', '00005');
INSERT INTO `street` VALUES ('63', null, null, null, '杨和', '000059', '00005');
INSERT INTO `street` VALUES ('64', null, null, null, '荷城', '000060', '00005');
INSERT INTO `street` VALUES ('65', null, null, null, '杨梅', '000061', '00005');
INSERT INTO `street` VALUES ('66', null, null, null, '白水井', '000062', '00006');
INSERT INTO `street` VALUES ('67', null, null, null, '逢源商业街', '000063', '00006');
INSERT INTO `street` VALUES ('68', null, null, null, '厚兴', '000064', '00006');
INSERT INTO `street` VALUES ('69', null, null, null, '湖滨北路', '000065', '00006');
INSERT INTO `street` VALUES ('70', null, null, null, '假日广场', '000066', '00006');
INSERT INTO `street` VALUES ('71', null, null, null, '莲塘路', '000067', '00006');
INSERT INTO `street` VALUES ('72', null, null, null, '岐关西路', '000068', '00006');
INSERT INTO `street` VALUES ('73', null, null, null, '石岐大信', '000069', '00006');
INSERT INTO `street` VALUES ('74', null, null, null, '孙文步行街', '000070', '00006');
INSERT INTO `street` VALUES ('75', null, null, null, '兴中广场', '000071', '00006');
INSERT INTO `street` VALUES ('76', null, null, null, '悦来路', '000072', '00006');
INSERT INTO `street` VALUES ('77', null, null, null, '中山北站', '000073', '00006');
INSERT INTO `street` VALUES ('78', null, null, null, '大兴路', '000074', '00007');
INSERT INTO `street` VALUES ('79', null, null, null, '南坦路', '000075', '00007');
INSERT INTO `street` VALUES ('80', null, null, null, '界狮南路', '000076', '00007');
INSERT INTO `street` VALUES ('81', null, null, null, '界狮北路', '000077', '00007');
INSERT INTO `street` VALUES ('82', null, null, null, '坦神南路', '000078', '00007');
INSERT INTO `street` VALUES ('83', null, null, null, '坦神北路', '000079', '00007');
INSERT INTO `street` VALUES ('84', null, null, null, '乐怡路', '000080', '00007');
INSERT INTO `street` VALUES ('85', null, null, null, '张家边', '000081', '00008');
INSERT INTO `street` VALUES ('86', null, null, null, '凯茵新城', '000082', '00008');
INSERT INTO `street` VALUES ('87', null, null, null, '濠头', '000083', '00008');
INSERT INTO `street` VALUES ('88', null, null, null, '中山港', '000084', '00008');
INSERT INTO `street` VALUES ('89', null, null, null, '沙岗墟', '000085', '00009');
INSERT INTO `street` VALUES ('90', null, null, null, '中山海关', '000086', '00009');
INSERT INTO `street` VALUES ('91', null, null, null, '远洋城', '000087', '00009');
INSERT INTO `street` VALUES ('92', null, null, null, '库充', '000088', '00009');
INSERT INTO `street` VALUES ('93', null, null, null, '兴中体育场', '000089', '00009');
INSERT INTO `street` VALUES ('94', null, null, null, '松苑路', '000090', '00009');
INSERT INTO `street` VALUES ('95', null, null, null, '槎桥洛', '000091', '00009');
INSERT INTO `street` VALUES ('96', null, null, null, '富华总站', '000092', '00010');
INSERT INTO `street` VALUES ('97', null, null, null, '富华道', '000093', '00010');
INSERT INTO `street` VALUES ('98', null, null, null, '沙朗', '000094', '00010');
INSERT INTO `street` VALUES ('99', null, null, null, '彩虹大道', '000095', '00010');
INSERT INTO `street` VALUES ('100', null, null, null, '菊城大道', '000096', '00011');
INSERT INTO `street` VALUES ('101', null, null, null, '文昌路', '000097', '00012');
INSERT INTO `street` VALUES ('102', null, null, null, '雅居乐', '000098', '00012');
INSERT INTO `street` VALUES ('103', null, null, null, '小琅环路', '000099', '00012');
INSERT INTO `street` VALUES ('104', null, null, null, '星宝路', '000100', '00013');
INSERT INTO `street` VALUES ('105', null, null, null, '龙瑞', '000101', '00013');
INSERT INTO `street` VALUES ('106', null, null, null, '庞头', '000102', '00013');
INSERT INTO `street` VALUES ('107', null, null, null, '康乐路', '000103', '00013');
INSERT INTO `street` VALUES ('108', null, null, null, '国贸', '000104', '00014');
INSERT INTO `street` VALUES ('109', null, null, null, '城南路', '000105', '00015');
INSERT INTO `street` VALUES ('110', null, null, null, '恒海路', '000106', '00015');
INSERT INTO `street` VALUES ('111', null, null, null, '北溪路', '000107', '00015');
INSERT INTO `street` VALUES ('112', null, null, null, '星晨花园', '000108', '00016');
INSERT INTO `street` VALUES ('113', null, null, null, '美景东路', '000109', '00016');
INSERT INTO `street` VALUES ('114', null, null, null, '美景西路', '000110', '00016');
INSERT INTO `street` VALUES ('115', null, null, null, '东港大道', '000111', '00017');
INSERT INTO `street` VALUES ('116', null, null, null, '龙起路', '000112', '00018');
INSERT INTO `street` VALUES ('117', null, null, null, '南朗镇政府', '000113', '00018');
INSERT INTO `street` VALUES ('118', null, null, null, '翠亨新区', '000114', '00018');
INSERT INTO `street` VALUES ('119', null, null, null, '凤翔大道', '000115', '00019');
INSERT INTO `street` VALUES ('120', null, null, null, '长安南路', '000116', '00020');
INSERT INTO `street` VALUES ('121', null, null, null, '南头大道', '000117', '00021');
INSERT INTO `street` VALUES ('122', null, null, null, '兴圃大道', '000118', '00022');
INSERT INTO `street` VALUES ('123', null, null, null, '金三大道', '000119', '00023');
INSERT INTO `street` VALUES ('124', null, null, null, '板芙北路', '000120', '00024');
INSERT INTO `street` VALUES ('125', null, null, null, '广药', '000121', '00025');
INSERT INTO `street` VALUES ('126', null, null, null, '理工', '000122', '00025');
INSERT INTO `street` VALUES ('127', null, null, null, '大涌镇政府', '000123', '00026');
INSERT INTO `street` VALUES ('128', null, null, null, '民众大道', '000124', '00027');
INSERT INTO `street` VALUES ('129', null, null, null, '埠港东路', '000125', '00028');
INSERT INTO `street` VALUES ('130', null, null, null, '神湾大道', '000126', '00029');
