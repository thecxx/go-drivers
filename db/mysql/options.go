package mysql

import (
	"time"
)

type ClientOptionHandlerFunc func(conf *Config)

// MaxConnLifetime
func WithMaxConnLifetime(lifetime time.Duration) ClientOptionHandlerFunc {
	return func(conf *Config) {
		conf.MaxLifetime = lifetime
	}
}

// MaxOpenConns
func WithMaxOpenConns(limit int) ClientOptionHandlerFunc {
	return func(conf *Config) {
		conf.MaxOpenConns = limit
	}
}

// MaxIdleConns
func WithMaxIdleConns(limit int) ClientOptionHandlerFunc {
	return func(conf *Config) {
		conf.MaxIdleConns = limit
	}
}

// DialTimeout
func WithDialTimeout(timeout time.Duration) ClientOptionHandlerFunc {
	return func(conf *Config) {
		conf.Timeout = timeout
	}
}

// ReadTimeout
func WithReadTimeout(timeout time.Duration) ClientOptionHandlerFunc {
	return func(conf *Config) {
		conf.ReadTimeout = timeout
	}
}

// WriteTimeout
func WithWriteTimeout(timeout time.Duration) ClientOptionHandlerFunc {
	return func(conf *Config) {
		conf.WriteTimeout = timeout
	}
}
