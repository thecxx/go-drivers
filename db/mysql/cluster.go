package mysql

import (
	"context"
	"errors"
	atomic2 "sync/atomic"
)

type Cluster struct {
	writers []*Database
	readers []*Database
	readerp int32
}

// New cluster with specific databases.
func NewCluster(writer *Database, readers ...*Database) (*Cluster, error) {
	if writer == nil {
		return nil, errors.New("invalid database for write")
	}
	if len(readers) > 0 {
		for _, reader := range readers {
			if reader == nil {
				return nil, errors.New("invalid database for read")
			}
		}
	} else {
		readers = make([]*Database, 0)
	}
	c := new(Cluster)
	c.writers = append(c.writers, writer)
	c.readers = append(c.readers, readers...)
	c.readerp = 0

	return c, nil
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (c *Cluster) Query(ctx context.Context, query string, args ...interface{}) (Result, error) {
	return c.reader().Query(ctx, query, args...)
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (c *Cluster) Exec(ctx context.Context, query string, args ...interface{}) (Result, error) {
	return c.writer().Exec(ctx, query, args...)
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
func (c *Cluster) BeginTransaction(ctx context.Context) (*Transaction, error) {
	return c.writer().BeginTransaction(ctx)
}

// Close stop the cluster.
func (c *Cluster) Close() {
	for _, writer := range c.writers {
		writer.Close()
	}
	for _, reader := range c.readers {
		reader.Close()
	}
}

// Get a database for write.
func (c *Cluster) writer() *Database {
	return c.writers[0]
}

// Get a database for read.
func (c *Cluster) reader() *Database {
	n := len(c.readers)
	switch {
	// 1. If no reader
	case n <= 0:
		return c.writer()
	// 2. Only one
	case n == 1:
		return c.readers[0]
	}
	// 3. Schedule
	return c.readers[atomic2.AddInt32(&c.readerp, 1)%int32(n)]
}
