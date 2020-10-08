package redis

import (
	"context"
	"time"

	driver "github.com/go-redis/redis/v8"
)

const (
	DefaultMinIdleConns = 5
	DefaultMaxLifetime  = 30 * time.Second
	DefaultDialTimeout  = 2 * time.Second
	DefaultReadTimeout  = 0 * time.Second
	DefaultWriteTimeout = 0 * time.Second
)

type Config struct {
	*driver.Options
}

// New a default config.
func NewDefaultConfig() *Config {
	c := new(Config)
	c.Options = new(driver.Options)

	c.Network = "tcp"
	c.DialTimeout = DefaultDialTimeout
	c.ReadTimeout = DefaultReadTimeout
	c.WriteTimeout = DefaultWriteTimeout
	c.MinIdleConns = DefaultMinIdleConns
	c.MaxConnAge = DefaultMaxLifetime

	return c
}

type Client struct {
	conf   *Config
	client *driver.Client
}

// New client.
func NewClient(addr, passwd string, options ...ClientOptionHandlerFunc) (*Client, error) {
	conf := NewDefaultConfig()
	conf.Addr = addr
	conf.Password = passwd
	// Apply options
	if len(options) > 0 {
		for _, handler := range options {
			handler(conf)
		}
	}
	return NewClientWithConfig(conf)
}

// New client with a specific config.
func NewClientWithConfig(conf *Config) (*Client, error) {
	// New a client with specific config.
	client := driver.NewClient(conf.Options)
	// Test ping
	err := client.Ping(context.Background()).Err()
	if err != nil {
		client.Close()
		return nil, err
	}
	return &Client{conf, client}, nil
}

// Ping verifies a connection to the cache server is still alive,
// establishing a connection if necessary.
func (c *Client) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

// Close closes the client, releasing any open resources.
//
// It is rare to Close a Client, as the Client is meant to be
// long-lived and shared between many goroutines.
func (c *Client) Close() error {
	return c.client.Close()
}
