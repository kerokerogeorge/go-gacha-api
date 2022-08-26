CREATE DATABASE IF NOT EXISTS go_test_dev;
USE go_test_dev;

CREATE TABLE IF NOT EXISTS users
(
    id            bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    name          varchar(36) DEFAULT NULL COMMENT 'ユーザー名',
    token         varchar(255) DEFAULT NULL COMMENT '認証トークン',
    created_at    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    updated_at    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY (id)
);

-- 3
CREATE TABLE IF NOT EXISTS characters
(
    id             bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    -- gacha_id       bigint(20) NOT NULL COMMENT 'ガチャID',
    name           varchar(36) DEFAULT NULL COMMENT 'キャラクター名',
    -- emission_rate  bigint(20) DEFAULT NULL COMMENT '排出確率',
    created_at     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    updated_at     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_characters
(
    id            bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    user_id       bigint(20) NOT NULL COMMENT 'ユーザーID',
    character_id  bigint(20) NOT NULL COMMENT 'キャラクターID',
    created_at    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    updated_at    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY(id),
    FOREIGN KEY(user_id) REFERENCES users (id),
    FOREIGN KEY(character_id) REFERENCES characters (id)
);

-- 1
CREATE TABLE IF NOT EXISTS gachas
(
    id             varchar(255) NOT NULL COMMENT 'id',
    created_at     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    updated_at     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY (id)
);

-- 2
CREATE TABLE IF NOT EXISTS character_emmition_rates
(
    id            bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    gacha_id      varchar(255) NOT NULL COMMENT 'ガチャID',
    character_id  bigint(20) NOT NULL COMMENT 'キャラクターID',
    emission_rate bigint(20) DEFAULT NULL COMMENT '排出確率',
    created_at    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    updated_at    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY(id),
    FOREIGN KEY(gacha_id) REFERENCES gachas (id),
    FOREIGN KEY(character_id) REFERENCES characters (id)
);

-- INSERT INTO characters (name, emission_rate, created_at, updated_at) VALUES ('pikachu', 10, NOW() , NOW());
-- INSERT INTO characters (name, emission_rate, created_at, updated_at) VALUES ('kimori', 20, NOW() , NOW());
-- INSERT INTO characters (name, emission_rate, created_at, updated_at) VALUES ('achamo', 30, NOW() , NOW());
-- INSERT INTO characters (name, emission_rate, created_at, updated_at) VALUES ('mizugoro', 40, NOW() , NOW());
