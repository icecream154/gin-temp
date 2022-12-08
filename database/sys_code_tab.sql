CREATE DATABASE /*!32312 IF NOT EXISTS */`db_goskeleton` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `db_goskeleton`;

/*Table structure for table `sys_code_tab` */

DROP TABLE IF EXISTS `sys_code_tab`;

CREATE TABLE `sys_code_tab`
(
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `phone`      VARCHAR(50) COMMENT '手机',
    `code`       VARCHAR(6) COMMENT '验证码',
    `issue_time` DATETIME COMMENT '签发时间',
    `checked`    TINYINT DEFAULT 0 COMMENT '0未校验 1已校验',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

ALTER TABLE `sys_code_tab` ADD INDEX (`phone`);