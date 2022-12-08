CREATE
    DATABASE /*!32312 IF NOT EXISTS */`db_goskeleton` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE
    `db_goskeleton`;

/*Table structure for table `sys_opinion_tab` */

DROP TABLE IF EXISTS `sys_opinion_tab`;

CREATE TABLE `sys_opinion_tab`
(
    `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `account`       VARCHAR(255) COMMENT '账号',
    `content`       VARCHAR(1000) DEFAULT '' COMMENT '意见内容',
    `image`         VARCHAR(300)  DEFAULT '' COMMENT '图片url',
    `contact`       VARCHAR(255) COMMENT '联系方式',
    `dealt`         TINYINT       DEFAULT 1 COMMENT '是否处理(1未处理，2已处理)',
    `handler`       VARCHAR(50)   DEFAULT '' COMMENT '处理人',
    `handle_result` VARCHAR(500)  DEFAULT '' COMMENT '处理结果',
    `delete`        TINYINT       DEFAULT 1 COMMENT '是否删除(1未删除，2已删除)',
    `created_at`    DATETIME      DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    DATETIME      DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;