package redis

type ClientOptionHandlerFunc func(conf *Config)

// DatabaseID
func WithDatabaseID(db int) ClientOptionHandlerFunc {
	return func(conf *Config) {
		conf.DB = db
	}
}
