package v8

import (
	"context"

	"github.com/go-redis/redis/v8"
	redislock "github.com/noahlsl/go-redislock"
)

type RedisAdapter struct {
	client redis.Cmdable
}

func New(client redis.Cmdable) redislock.RedisInter {
	return &RedisAdapter{client: client}
}

func (r *RedisAdapter) Eval(ctx context.Context, script string, keys []string, args ...interface{}) redislock.RedisCmd {
	cmd := r.client.Eval(ctx, script, keys, args...)
	return &RedisCmdWrapper{cmd: cmd}
}

type RedisCmdWrapper struct {
	cmd *redis.Cmd
}

func (w *RedisCmdWrapper) Result() (interface{}, error) {
	return w.cmd.Result()
}
func (w *RedisCmdWrapper) Int64() (int64, error) {
	return w.cmd.Int64()
}
