package main

import (
	"encoding/json"
)

type Config struct {
	ListenAddress string `config:"LISTEN_ADDRESS"`
	MySqlUsername string `config:"MYSQL_USERNAME"`
	MySqlPassword string `config:"MYSQL_PASSWORD"`
	MySqlDbName   string `config:"MYSQL_DBNAME"`
	MySqlAddress  string `config:"MYSQL_ADDRESS"`
	BaseUrl       string `config:"BASE_URL"`
}

func NewConfigDefault() *Config {
	return &Config{
		ListenAddress: "127.0.0.1:8080",
		MySqlUsername: "user",
		MySqlPassword: "123456",
		MySqlDbName:   "multifinance-api",
		MySqlAddress:  "127.0.0.1:3306",
		BaseUrl:       "http://localhost/",
	}
}

func (c *Config) AsString() string {
	data, _ := json.Marshal(c)
	return string(data)
}
