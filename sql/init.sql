CREATE USER 'admin'@'localhost';
GRANT ALL ON *.* TO 'admin'@'localhost';

CREATE DATABASE ferry;

USE ferry;

CREATE TABLE user (
    id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    score       INT UNSIGNED NOT NULL,
    name        TEXT NOT NULL,
    timestamp   TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE pass (
    id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id     TEXT NOT NULL,
    point_id    TEXT NOT NULL,
    timestamp   TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
