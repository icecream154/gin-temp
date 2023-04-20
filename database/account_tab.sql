CREATE DATABASE /*!32312 IF NOT EXISTS */`youni_intelligence` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `youni_intelligence`;

/*Table structure for table `account_tab` */

DROP TABLE IF EXISTS `account_tab`;

CREATE TABLE `account_tab`
(
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `phone`      VARCHAR(50) COMMENT '手机',
    `email`      VARCHAR(150) COMMENT '邮箱',
    `password`   VARCHAR(256) COMMENT '密码',
    `code`       VARCHAR(12) COMMENT '邀请码',
    `created_at` DATETIME COMMENT '签发时间' DEFAULT NOW(),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

ALTER TABLE `account_tab`
    ADD INDEX (`phone`);