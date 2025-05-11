-- 데이터베이스 생성 (기존과 동일)
CREATE DATABASE IF NOT EXISTS coupon_issuance_system;
USE coupon_issuance_system;

-- campaign 테이블
CREATE TABLE IF NOT EXISTS `campaign` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    `available_coupons` INT UNSIGNED NOT NULL DEFAULT 0,
    `start_date_time` DATETIME(6) NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- coupon 테이블
CREATE TABLE IF NOT EXISTS `coupon` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `campaign_id` INT UNSIGNED NOT NULL,
    `code` VARCHAR(10) NOT NULL UNIQUE,
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`campaign_id`) REFERENCES `campaign`(`id`) ON DELETE CASCADE,
    KEY `idx_campaign_id` (`campaign_id`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;