package mysql

import (
	"context"
	"testing"
)

// New database.
func newClient() (*Client, error) {
	return NewClient("127.0.0.1:3306", "test", "root", "123456")
}

func TestClient_Query(t *testing.T) {
	d, err := newClient()
	if err != nil {
		t.Errorf("New database failed, err=%s\n", err.Error())
		return
	}
	result, err := d.Query(context.Background(), "SELECT * FROM `test` LIMIT 1")
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
