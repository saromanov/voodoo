package redis

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/saromanov/voodoo/pkg/source"
)

type Redis struct {
	ctx    context.Context
	client *redis.Client
	out    chan interface{}
}

func New(ctx context.Context, config *redis.Options) source.Source {
	return &Redis{
		ctx:    ctx,
		client: redis.NewClient(config),
		out:    make(chan interface{}),
	}
}

func (r *Redis) init(ch <-chan *redis.Message) {
	for {
		select {
		case <-r.ctx.Done():
			break
		case msg := <-ch:
			r.out <- msg
		}
	}

	close(r.out)
	r.client.Close()
}

func (r *Redis) With() error {
	return nil
}

func (r *Redis) To() error {
	return nil
}
