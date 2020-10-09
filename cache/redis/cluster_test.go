package redis

import (
	"context"
	"testing"
)

// New cluster.
func newCluster() (*Cluster, error) {
	return NewCluster([]string{"127.0.0.1:6379"}, "")
}

func TestCluster_Hash(t *testing.T) {
	cluster, err := newCluster()
	if err != nil {
		t.Errorf("New client failed, err=%s\n", err.Error())
		return
	}

	cluster.HSet(context.Background(), "unittest_hash", "unit1", "yes1")
	cluster.HSet(context.Background(), "unittest_hash", "unit2", "yes2")
	cluster.HSet(context.Background(), "unittest_hash", "unit3", "yes3")

	hash, err := cluster.Hash(context.Background(), "unittest_hash", "unit", "unit1")
	if err != nil {
		t.Errorf("Execute command failed, err=%s\n", err.Error())
		return
	}
	t.Logf("%#v\n", hash)
}
