CREATE TABLE `car`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'id_inc',
  `car_rfid_id` bigint(0) NULL COMMENT '车辆标识',
  `marster_id` bigint(0) NULL COMMENT '车主ID',
  `type` varchar(20) NULL COMMENT '车辆类型',
  `plate_number` varchar(20) NULL COMMENT '车牌号码',
  `color` varchar(20) NULL COMMENT '车辆颜色',
  `start_time` datetime(0) NULL COMMENT '开始时间',
  `flag` varchar(2) NULL COMMENT '使用标志',
  PRIMARY KEY (`id`)
);

CREATE TABLE `manager`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'id_inc',
  `nickname` varchar(20) NULL COMMENT '用户名',
  `password` varchar(20) NULL COMMENT '密码',
  PRIMARY KEY (`id`)
);

CREATE TABLE `master`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'id_inc',
  `id_number` varchar(20) NULL COMMENT '身份证号',
  `name` varchar(20) NULL COMMENT '车主姓名',
  `nickname` varchar(20) NULL COMMENT '车主名称',
  `open_id` varchar(45) NULL COMMENT '小程序用户ID',
  `mobile_number` varchar(13) NULL COMMENT '电话号码',
  `register_time` datetime(0) NULL COMMENT '注册时间',
  PRIMARY KEY (`id`)
);

CREATE TABLE `position_desc`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'id_inc',
  `rfid_id` bigint(0) NULL COMMENT '基站ID',
  `position` varchar(255) NULL COMMENT '地理位置',
  PRIMARY KEY (`id`)
);

CREATE TABLE `rfid_data`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'id_inc',
  `rfid_id` bigint(0) NULL COMMENT '基站ID',
  `car_rfid_id` bigint(0) NULL COMMENT '机动车标识',
  `time` datetime(0) NULL COMMENT '记录时间',
  PRIMARY KEY (`id`)
);

