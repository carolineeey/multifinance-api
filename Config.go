package api

import (
	"encoding/json"
)

type Config struct {
	ListenAddress    []string `config:"LISTEN_ADDRESS"`
	EnableTls        bool     `config:"ENABLE_TLS"`
	TlsListenAddress []string `config:"TLS_LISTEN_ADDRESS"`
	TlsCertFile      string   `config:"TLS_CERT_FILE"`
	TlsKeyFile       string   `config:"TLS_KEY_FILE"`
	MySqlUsername    string   `config:"MYSQL_USERNAME"`
	MySqlPassword    string   `config:"MYSQL_PASSWORD"`
	MySqlDbName      string   `config:"MYSQL_DBNAME"`
	MySqlAddress     string   `config:"MYSQL_ADDRESS"`
	BaseUrl          string   `config:"BASE_URL"`
}

func NewConfigDefault() *Config {
	return &Config{
		ListenAddress:    []string{"127.0.0.1:8080"},
		TlsListenAddress: []string{"127.0.0.1:443"},
		TlsCertFile:      "certs/localhost.crt",
		TlsKeyFile:       "certs/localhost.key",
		MySqlUsername:    "root",
		MySqlPassword:    "123456",
		MySqlDbName:      "users",
		MySqlAddress:     "tcp(127.0.0.1:3306)",
		BaseUrl:          "http://localhost/",
	}
}

func (c *Config) AsString() string {
	data, _ := json.Marshal(c)
	return string(data)
}
