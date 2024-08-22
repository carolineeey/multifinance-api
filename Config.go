package api

import (
	"encoding/json"
)

type Config struct {
	ListenAddress           []string `config:"LISTEN_ADDRESS"`
	EnableTls               bool     `config:"ENABLE_TLS"`
	TlsListenAddress        []string `config:"TLS_LISTEN_ADDRESS"`
	TlsCertFile             string   `config:"TLS_CERT_FILE"`
	TlsKeyFile              string   `config:"TLS_KEY_FILE"`
	MySqlUsername           string   `config:"MYSQL_USERNAME"`
	MySqlPassword           string   `config:"MYSQL_PASSWORD"`
	MySqlDbName             string   `config:"MYSQL_DBNAME"`
	MySqlAddress            string   `config:"MYSQL_ADDRESS"`
	PrometheusListenAddress string   `config:"PROMETHEUS_LISTEN_ADDRESS"`
	LoggingToStd            bool     `config:"LOGGING_TO_STD"`
	LoggingStdColor         bool     `config:"LOGGING_STD_COLOR"`
	LoggingToFile           bool     `config:"LOGGING_TO_FILE"`
	LoggingToFluentd        bool     `config:"LOGGING_TO_FLUENTD"`
	LoggingFilePath         string   `config:"LOGGING_FILE_PATH"`
	LoggingFluentdHost      string   `config:"LOGGING_FLUENTD_HOST"`
	LoggingFluentdPort      int      `config:"LOGGING_FLUENTD_PORT"`
	LoggingFluentdPath      string   `config:"LOGGING_FLUENTD_PATH"`
	LoggingVerbosity        int      `config:"LOGGING_VERBOSITY"`
	BaseUrl                 string   `config:"BASE_URL"`
}

func NewConfigDefault() *Config {
	return &Config{
		ListenAddress:           []string{"127.0.0.1:8080"},
		TlsListenAddress:        []string{"127.0.0.1:443"},
		TlsCertFile:             "certs/localhost.crt",
		TlsKeyFile:              "certs/localhost.key",
		MySqlUsername:           "root",
		MySqlPassword:           "123456",
		MySqlDbName:             "multifinance-api",
		MySqlAddress:            "tcp(127.0.0.1:3306)",
		PrometheusListenAddress: "0.0.0.0:8081",
		LoggingToStd:            true,
		LoggingStdColor:         true,
		LoggingToFile:           false,
		BaseUrl:                 "http://localhost/",
	}
}

func (c *Config) AsString() string {
	data, _ := json.Marshal(c)
	return string(data)
}
