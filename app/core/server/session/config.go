package session

import (
	"log"
	"os"
)

type Listen struct {
	Port   string `yaml:"port"`
	BindIP string `yaml:"bind_ip"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:""`
	SSLMode  string
}

type Config struct {
	DB     DB     `yaml:"db"`
	Listen Listen `yaml:"listen"`
}

func GetDBConfig() *DB {
	return &DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSL_MODE"),
	}
}

func GetListenConfig() *Listen {
	return &Listen{
		Port:   os.Getenv("LISTEN_PORT"),
		BindIP: os.Getenv("BIND_IP"),
	}
}

func GetConfig() *Config {
	config := &Config{}
	config.DB = *GetDBConfig()
	config.Listen = *GetListenConfig()

	log.Println(config)

	return config
}
