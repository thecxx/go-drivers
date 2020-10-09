package redis

import (
	"context"
	"errors"

	driver "github.com/go-redis/redis/v8"
)

type ClusterConfig struct {
	*driver.ClusterOptions
}

// New default cluster config.
func NewDefaultClusterConfig() *ClusterConfig {
	c := new(ClusterConfig)
	c.ClusterOptions = new(driver.ClusterOptions)

	c.DialTimeout = DefaultDialTimeout
	c.ReadTimeout = DefaultReadTimeout
	c.WriteTimeout = DefaultWriteTimeout
	c.MinIdleConns = DefaultMinIdleConns
	c.MaxConnAge = DefaultMaxLifetime

	return c
}

type Cluster struct {
	*handlers
	conf   *ClusterConfig
	client *driver.ClusterClient
}

// New cluster.
func NewCluster(addrs []string, password string, options ...ClusterOptionHandlerFunc) (*Cluster, error) {
	// Check address list
	if len(addrs) <= 0 {
		return nil, errors.New("no cluster address found")
	} else {
		for _, addr := range addrs {
			if addr == "" {
				return nil, errors.New("invalid cluster address")
			}
		}
	}
	conf := NewDefaultClusterConfig()
	conf.Addrs = addrs
	conf.Password = password
	// Apply options
	if len(options) > 0 {
		for _, handler := range options {
			handler(conf)
		}
	}
	return NewClusterWithConfig(conf)
}

// New cluster client with a specific config.
func NewClusterWithConfig(conf *ClusterConfig) (*Cluster, error) {
	// New a cluster client with specific config.
	client := driver.NewClusterClient(conf.ClusterOptions)
	// Test ping
	err := client.Ping(context.Background()).Err()
	if err != nil {
		client.Close()
		return nil, err
	}
	c := &Cluster{
		conf:   conf,
		client: client,
	}
	c.handlers = &handlers{client}

	return c, nil
}

// Close stop the cluster.
func (c *Cluster) Close() error {
	return c.client.Close()
}
