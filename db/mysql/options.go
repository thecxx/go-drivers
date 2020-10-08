package mysql

import (
	"time"
)

type DatabaseOptionHandlerFunc func(conf *Config)

// MaxConnLifetime
func WithMaxConnLifetime(lifetime time.Duration) DatabaseOptionHandlerFunc {
	return func(conf *Config) {
		conf.MaxLifetime = lifetime
	}
}

// MaxOpenConns
func WithMaxOpenConns(limit int) DatabaseOptionHandlerFunc {
	return func(conf *Config) {
		conf.MaxOpenConns = limit
	}
}

// MaxIdleConns
func WithMaxIdleConns(limit int) DatabaseOptionHandlerFunc {
	return func(conf *Config) {
		conf.MaxIdleConns = limit
	}
}

// DialTimeout
func WithDialTimeout(timeout time.Duration) DatabaseOptionHandlerFunc {
	return func(conf *Config) {
		conf.Timeout = timeout
	}
}

// ReadTimeout
func WithReadTimeout(timeout time.Duration) DatabaseOptionHandlerFunc {
	return func(conf *Config) {
		conf.ReadTimeout = timeout
	}
}

// WriteTimeout
func WithWriteTimeout(timeout time.Duration) DatabaseOptionHandlerFunc {
	return func(conf *Config) {
		conf.WriteTimeout = timeout
	}
}
