package api

import (
	"github.com/go-resty/resty/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sony/gobreaker"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var multifinanceClient *api.Client

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // Flushes buffer, if any

	globalConfig := NewConfigDefault()

	viper.SetConfigFile("config.yaml") // You can adjust the configuration file format and path as needed
	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal("cannot parse configuration", zap.Error(err))
		os.Exit(1)
		return
	}

	mustLoadConfig(globalConfig)
	logger.Info("using configuration", zap.String("config", globalConfig.AsString()))

	cbSettings := gobreaker.Settings{
		Name:    "MULTIFINANCE_API",
		Timeout: 10 * time.Second,
	}
	rcb := gobreaker.NewCircuitBreaker(cbSettings)

	// Initialize Client
	multifinanceClient = api.NewDefaultClient(globalConfig.BaseUrl, rcb, logger)

	// Spawn the Prometheus metric server
	http.Handle("/metrics", promhttp.Handler())

	// Catch SIGTERM signal and shut down the system gracefully.
	// If the CB is tripped, also throw exit code 1 and add a 30 seconds timer to shut down the system forcefully.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		err := http.ListenAndServe(globalConfig.PrometheusListenAddress, nil)
		if err != nil {
			logger.Fatal("Prometheus server failed to start", zap.Error(err))
		}
	}()

	serverManager := resty.New() // Using resty as a basic HTTP client manager
	serverManager.SetLogger(logger.Sugar())

	router := createRouter(logger)
	httpServer := &http.Server{
		Addr:    globalConfig.ListenAddress,
		Handler: router,
	}

	// Catch SIGTERM signal and shut down the system gracefully.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		logger.Warn("shutting down")
		httpServer.Close()
	}()

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("server failed to start", zap.Error(err))
	}
}

// Loads the configuration, and terminates the process if configuration cannot be loaded.
func mustLoadConfig(cfg *Config) {
	if err := viper.Unmarshal(cfg); err != nil {
		logger.Fatal("cannot parse configuration", zap.Error(err))
		os.Exit(1)
	}

	logger.Info("configuration loaded", zap.Any("config", *cfg))
	return
}
