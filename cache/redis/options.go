package redis

type ClientOptionHandlerFunc func(conf *Config)

// Username
func WithUsername(username string) ClientOptionHandlerFunc {
	return func(conf *Config) {
		conf.Username = username
	}
}

// DatabaseID
func WithDatabaseID(db int) ClientOptionHandlerFunc {
	return func(conf *Config) {
		conf.DB = db
	}
}

type ClusterOptionHandlerFunc func(conf *ClusterConfig)

// Cluster username
func WithClusterUsername(username string) ClusterOptionHandlerFunc {
	return func(conf *ClusterConfig) {
		conf.Username = username
	}
}
