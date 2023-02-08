CREATE DATABASE IF NOT EXISTS `url_shortner`;
USE `url_shortner`

CREATE TABLE Link (
    Id int NOT NULL AUTO_INCREMENET,
    Domain varchar(20) NOT_NULL,
    Url varchar(50) NOT_NULL,
    ShortUrl varchar(20) NOT_NULL,
    ShortId varchar(10) NOT_NULL PRIMARY KEY,
    Clicks int DEFAULT 0
    Created DATE GETDATE()
);