package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	driver "github.com/go-sql-driver/mysql"
)

const (
	DefaultMaxOpenConns = 50
	DefaultMaxIdleConns = 10
	DefaultMaxLifetime  = 30 * time.Second
	DefaultDialTimeout  = 2 * time.Second
	DefaultReadTimeout  = 0 * time.Second
	DefaultWriteTimeout = 0 * time.Second
)

type Config struct {
	*driver.Config
	// Extension
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  time.Duration
}

// New a default config.
func NewDefaultConfig() *Config {
	c := new(Config)
	c.Config = driver.NewConfig()

	c.Net = "tcp"
	c.Timeout = DefaultDialTimeout
	c.ReadTimeout = DefaultReadTimeout
	c.WriteTimeout = DefaultWriteTimeout
	c.MaxOpenConns = DefaultMaxOpenConns
	c.MaxIdleConns = DefaultMaxIdleConns
	c.MaxLifetime = DefaultMaxLifetime

	return c
}

// Generate an unique ID.
func (c *Config) UniqId() string {
	return fmt.Sprintf("%s://%s/%s", c.Net, c.Addr, c.DBName)
}

type Database struct {
	db   *sql.DB
	conf *Config
}

// New database.
func NewDatabase(addr, dbname, user, passwd string, options ...DatabaseOptionHandlerFunc) (*Database, error) {
	conf := NewDefaultConfig()
	conf.Addr = addr
	conf.DBName = dbname
	conf.User = user
	conf.Passwd = passwd
	// Apply options
	if len(options) > 0 {
		for _, handler := range options {
			handler(conf)
		}
	}
	return NewDatabaseWithConfig(conf)
}

// New database with a specific DSN.
func NewDatabaseWithDSN(dsn string, maxLifetime time.Duration, maxOpenConns, maxIdleConns int) (*Database, error) {
	// Parse DSN
	cfg, err := driver.ParseDSN(dsn)
	if err != nil {
		return nil, err
	}
	conf := &Config{
		Config:       cfg,
		MaxOpenConns: maxOpenConns,
		MaxIdleConns: maxIdleConns,
		MaxLifetime:  maxLifetime,
	}
	return NewDatabaseWithConfig(conf)
}

// New database with a specific config.
func NewDatabaseWithConfig(conf *Config) (*Database, error) {
	// Open a specific database
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, err
	}
	// Setting
	db.SetConnMaxLifetime(conf.MaxLifetime)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	// Test ping
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return &Database{db, conf}, nil
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (d *Database) Query(ctx context.Context, query string, args ...interface{}) (Result, error) {
	rows, err := d.db.QueryContext(ctx, query, args...)
	return Result{rows: rows}, err
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (d *Database) Exec(ctx context.Context, query string, args ...interface{}) (Result, error) {
	result, err := d.db.ExecContext(ctx, query, args...)
	return Result{result: result}, err
}

// BeginTransaction starts a transaction.
//
// The provided context is used until the transaction is committed or rolled back.
// If the context is canceled, the sql package will roll back
// the transaction. Tx.Commit will return an error if the context provided to
// BeginTx is canceled.
//
// The provided TxOptions is optional and may be nil if defaults should be used.
// If a non-default isolation level is used that the driver doesn't support,
// an error will be returned.
func (d *Database) BeginTransaction(ctx context.Context) (*Transaction, error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx}, nil
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (d *Database) Ping(ctx context.Context) error {
	return d.db.PingContext(ctx)
}

// Get the number of connections currently in use.
func (d *Database) ActiveConns() int {
	return d.db.Stats().InUse
}

// Get the number of idle connections.
func (d *Database) IdleConns() int {
	return d.db.Stats().Idle
}

// Close closes the database and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server
// to finish.
//
// It is rare to Close a DB, as the DB handle is meant to be
// long-lived and shared between many goroutines.
func (d *Database) Close() error {
	return d.db.Close()
}
