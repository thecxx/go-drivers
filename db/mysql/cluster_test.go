package mysql

import (
	"context"
	"testing"
)

// New cluster.
func newCluster(writer *Client, readers ...*Client) (*Cluster, error) {
	return NewCluster(writer, readers...)
}

func TestCluster_Query(t *testing.T) {
	d1, err := newClient()
	if err != nil {
		t.Errorf("New database failed, err=%s\n", err.Error())
		return
	}
	c, err := newCluster(d1)
	if err != nil {
		t.Errorf("New cluster failed, err=%s\n", err.Error())
		return
	}
	result, err := c.Query(context.Background(), "SELECT * FROM `test` LIMIT 1")
	if err != nil {
		t.Errorf("Query failed, err=%s\n", err.Error())
		return
	}
	rows, err := result.Rows()
	if err != nil {
		t.Errorf("Get rows failed, err=%s\n", err.Error())
		return
	}
	if rows == nil {
		t.Errorf("Invalid rows\n")
		return
	}
}
