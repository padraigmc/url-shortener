CREATE DATABASE IF NOT EXISTS `url_shortener`;

USE `url_shortener`;

CREATE TABLE IF NOT EXISTS Link (
    Id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    Domain varchar(20) NOT NULL,
    Url varchar(50) NOT NULL,
    ShortUrl varchar(20) NOT NULL,
    ShortId varchar(10) NOT NULL,
    Clicks int DEFAULT 0,
    Created DATETIME DEFAULT NOW()
);