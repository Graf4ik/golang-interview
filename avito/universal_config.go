package avito

import (
	"fmt"
	"time"
)

// Реализовать универсальный конфиг, который использует опциональный паттерн

type Config struct {
	Host         string
	Port         int
	Timeout      time.Duration
	EnableLogger bool
}

type Option func(*Config)

func main() {
	cfg := NewConfig(
		WithHost("127.0.0.1"),
		WithPort(9090),
		WithTimeout(10*time.Second),
		WithLogger(true),
	)

	fmt.Printf("%+v\n", cfg)
}

// Конструктор с применением опций
func NewConfig(opts ...Option) Config {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

func defaultConfig() Config {
	return Config{
		Host:         "localhost",
		Port:         8080,
		Timeout:      1 * time.Second,
		EnableLogger: true,
	}
}

func WithHost(host string) Option {
	return func(config *Config) {
		config.Host = host
	}
}

func WithPort(port int) Option {
	return func(config *Config) {
		config.Port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(config *Config) {
		config.Timeout = timeout
	}
}

func WithLogger(enableLogger bool) Option {
	return func(config *Config) {
		config.EnableLogger = enableLogger
	}
}
