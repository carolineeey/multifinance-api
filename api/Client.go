package api

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sony/gobreaker"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Client struct {
	URL               string
	Breaker           *gobreaker.CircuitBreaker
	Timeout           time.Duration
	Logger            *zap.Logger
	GenericApiMetrics *promhttp.InstrumentTrace
}

// NewClient creates a new client with a default logger.
func NewClient(
	url string,
	breaker *gobreaker.CircuitBreaker,
	timeout time.Duration,
	logger *zap.Logger,
) *Client {
	return &Client{
		URL:     url,
		Breaker: breaker,
		Timeout: timeout,
		Logger:  logger,
	}
}

// NewDefaultClient creates a default api.Client with timeout set to 30 seconds.
func NewDefaultClient(listenAddress string, breaker *gobreaker.CircuitBreaker, logger *zap.Logger) *Client {
	return NewClient(
		listenAddress,
		breaker,
		30*time.Second,
		logger,
	)
}

// genericApiMetricWrapper wraps a handler with GenericApiMetricsWrapper, only if GenericApiMetrics is not nil.
func (c *Client) genericApiMetricWrapper(handler http.Handler) http.Handler {
	if c.GenericApiMetrics == nil {
		return handler
	}

	return promhttp.InstrumentHandlerCounter(nil, handler, nil)
}
