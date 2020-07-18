package redis

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/saromanov/voodoo/pkg/source"
	"github.com/saromanov/voodoo/pkg/transform"
)

type Redis struct {
	ctx    context.Context
	client *redis.Client
	out    chan interface{}
}

// New creates redis source
func New(ctx context.Context, config *redis.Options) source.Source {
	if config == nil {
		config = &redis.Options{}
	}
	return &Redis{
		ctx:    ctx,
		client: redis.NewClient(config),
		out:    make(chan interface{}),
	}
}

// prepare converts input redis message to data interface
func (r *Redis) prepare(in interface{}) interface{} {
	result := in.(*redis.Message)
	return result.Payload
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

func (r *Redis) With(transform.Transform) source.Source {
	return r
}

func (r *Redis) To() error {
	return nil
}
