package redis

import (
	"context"
	"testing"
)

// New client.
func newClient() (*Client, error) {
	return NewClient("127.0.0.1:6379", "")
}

func TestClient_Set(t *testing.T) {
	client, err := newClient()
	if err != nil {
		t.Errorf("New client failed, err=%s\n", err.Error())
		return
	}
	ok, err := client.Set(context.Background(), "unittest_string", "yes", 0)
	if err != nil {
		t.Errorf("Execute command `SET` failed, err=%s\n", err.Error())
		return
	}
	if !ok {
		t.Errorf("EXEcute command `SET` failed, err=%s\n", err.Error())
		return
	}
}

func TestClient_HSet(t *testing.T) {
	client, err := newClient()
	if err != nil {
		t.Errorf("New client failed, err=%s\n", err.Error())
		return
	}
	ok, err := client.HSet(context.Background(), "unittest_hash", "unit", "yes", "unit8", "yes8")
	if err != nil {
		t.Errorf("Execute command `HSET` failed, err=%s\n", err.Error())
		return
	}
	if !ok {
		t.Errorf("Execute command `HSET` failed, err=%s\n", err.Error())
		return
	}
}

func TestClient_Hash(t *testing.T) {
	client, err := newClient()
	if err != nil {
		t.Errorf("New client failed, err=%s\n", err.Error())
		return
	}

	client.HSet(context.Background(), "unittest_hash", "unit1", "yes1")
	client.HSet(context.Background(), "unittest_hash", "unit2", "yes2")
	client.HSet(context.Background(), "unittest_hash", "unit3", "yes3")

	hash, err := client.Hash(context.Background(), "unittest_hash", "unit", "unit1")
	if err != nil {
		t.Errorf("Execute command failed, err=%s\n", err.Error())
		return
	}
	t.Logf("%#v\n", hash)
}

func TestClient_Type(t *testing.T) {
	client, err := newClient()
	if err != nil {
		t.Errorf("New client failed, err=%s\n", err.Error())
		return
	}

	_type, err := client.Type(context.Background(), "unittest_hash")
	if err != nil {
		t.Errorf("Execute command failed, err=%s\n", err.Error())
		return
	}
	t.Logf("%#v\n", _type)
}
