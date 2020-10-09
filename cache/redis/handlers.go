package redis

import (
	"context"
	"errors"
	"time"

	driver "github.com/go-redis/redis/v8"
)

type handlers struct {
	cmdable driver.Cmdable
}

// Server

// Ping verifies a connection to the cache server is still alive,
// establishing a connection if necessary.
func (h *handlers) Ping(ctx context.Context) error {
	return h.cmdable.Ping(ctx).Err()
}

// Common

// `TYPE key` command.
func (h *handlers) Type(ctx context.Context, key string) (string, error) {
	return h.cmdable.Type(ctx, key).Result()
}

// `DEL key` command.
func (h *handlers) Delete(ctx context.Context, keys ...string) error {
	return h.cmdable.Del(ctx, keys...).Err()
}

// `EXISTS key` command.
func (h *handlers) Exists(ctx context.Context, key string) (bool, error) {
	return h.assertInt(h.cmdable.Exists(ctx, key), 1)
}

// `PERSIST key` command.
func (h *handlers) Persist(ctx context.Context, key string) (bool, error) {
	return h.cmdable.Persist(ctx, key).Result()
}

// `TTL key time` command.
func (h *handlers) TTL(ctx context.Context, key string) (time.Duration, error) {
	return h.cmdable.TTL(ctx, key).Result()
}

// `EXPIRE key expire` command.
func (h *handlers) Expire(ctx context.Context, key string, expire time.Duration) (bool, error) {
	return h.cmdable.Expire(ctx, key, expire).Result()
}

// `EXPIREAT key time` command.
func (h *handlers) ExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	return h.cmdable.ExpireAt(ctx, key, tm).Result()
}

// String

// `GET key` command.
func (h *handlers) String(ctx context.Context, key string) (string, error) {
	return h.cmdable.Get(ctx, key).Result()
}

// `GET key` command.
func (h *handlers) Get(ctx context.Context, key string) (string, error) {
	return h.cmdable.Get(ctx, key).Result()
}

// `SET key value [expiration]` command.
func (h *handlers) Set(ctx context.Context, key string, value interface{}, expire time.Duration) (bool, error) {
	return h.assertOK(h.cmdable.Set(ctx, key, value, expire))
}

// `INCR key` command.
func (h *handlers) Incr(ctx context.Context, key string) (int64, error) {
	return h.cmdable.Incr(ctx, key).Result()
}

// `DECR key` command.
func (h *handlers) Decr(ctx context.Context, key string) (int64, error) {
	return h.cmdable.Decr(ctx, key).Result()
}

// `INCRBY key value` command.
func (h *handlers) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return h.cmdable.IncrBy(ctx, key, value).Result()
}

// `DECRBY key value` command.
func (h *handlers) DecrBy(ctx context.Context, key string, value int64) (int64, error) {
	return h.cmdable.DecrBy(ctx, key, value).Result()
}

// Hash

// `HGETALL key` command.
// `HMGET key field [field]`.
func (h *handlers) Hash(ctx context.Context, key string, fields ...string) (map[string]string, error) {
	n := len(fields)
	if n <= 0 {
		return h.cmdable.HGetAll(ctx, key).Result()
	}
	values, err := h.cmdable.HMGet(ctx, key, fields...).Result()
	if err != nil {
		return nil, err
	}
	result := make(map[string]string, n)
	for index, field := range fields {
		value, ok := values[index].(string)
		if ok {
			result[field] = value
		} else {
			result[field] = ""
		}
	}
	return result, nil
}

// `HEXISTS key field` command.
func (h *handlers) HExists(ctx context.Context, key, field string) (bool, error) {
	return h.cmdable.HExists(ctx, key, field).Result()
}

// `HLEN key` command.
func (h *handlers) HLen(ctx context.Context, key string) (int64, error) {
	return h.cmdable.HLen(ctx, key).Result()
}

// `HKEYS key` command.
func (h *handlers) HKeys(ctx context.Context, key string) ([]string, error) {
	return h.cmdable.HKeys(ctx, key).Result()
}

// `HVALS key` command.
func (h *handlers) HVals(ctx context.Context, key string) ([]string, error) {
	return h.cmdable.HVals(ctx, key).Result()
}

// `HGET key field` command.
func (h *handlers) HGet(ctx context.Context, key, field string) (string, error) {
	return h.cmdable.HGet(ctx, key, field).Result()
}

// `HSET key field value [field [value]]` command.
func (h *handlers) HSet(ctx context.Context, key string, args ...interface{}) (bool, error) {
	if n := len(args); n <= 0 || n%2 != 0 {
		return false, errors.New("invalid command args")
	}
	return h.cmdable.HMSet(ctx, key, args...).Result()
}

// List

// Set

// ZSet

func (h *handlers) assertOK(cmd *driver.StatusCmd) (bool, error) {
	value, err := cmd.Result()
	if err != nil {
		return false, err
	}
	return "OK" == value, nil
}

func (h *handlers) assertInt(cmd *driver.IntCmd, expect int64) (bool, error) {
	value, err := cmd.Result()
	if err != nil {
		return false, err
	}
	return expect == value, nil
}
