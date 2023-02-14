package config

import (
	"fmt"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "url-shortener-database-1.ch8t8lotsnfu.eu-west-1.rds.amazonaws.com",
			Port:     3306,
			Username: "admin",
			Password: "password",
			Name:     "url_shortener",
			Charset:  "utf8",
		},
	}
}

func (config *Config) GetDBUri() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)
}