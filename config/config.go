package config

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Server struct {
		Host	string		`yaml:"host", envconfig:"SERVER_HOST"`
		Port	string 		`yaml:"port", envconfig:"SERVER_PORT"`
	} `yaml:"server"`

	Database struct {
		Dialect  string		`yaml:"dialect"`
		Host     string		`yaml:"host", envconfig:"DB_HOST"`
		Port     int		`yaml:"port", envconfig:"DB_PORT"`
		Username string		`yaml:"username", envconfig:"DB_USERNAME"`
		Password string		`yaml:"password", envconfig:"DB_PASSWORD"`
		DBName   string		`yaml:"db_name"`
		Charset  string		`yaml:"charset"`
	} `yaml:"database"`
}

func NewConfig() *Config {
	config := &Config{}
	config.readFile()
	config.readEnv()
	return config
}

func (config *Config) readFile() {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Error(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Error(err)
	}
}

func (config *Config) readEnv() { 
    err := envconfig.Process("", config) 
    if err != nil { 
        log.Error(err)
    }
}

func (config *Config) GetDBUri() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
		config.Database.Charset)
}