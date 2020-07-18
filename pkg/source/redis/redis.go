package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/saromanov/voodoo/pkg/source"
	"github.com/saromanov/voodoo/pkg/transform"
)

// Options defines initialization options for redis source
type Options struct {
	ClientOptions *redis.Options
	Channel       string
}

type Redis struct {
	ctx    context.Context
	client *redis.Client
	config *Options
	out    chan interface{}
}

// New creates redis source
func New(ctx context.Context, config *Options) (source.Source, error) {
	if config == nil {
		return nil, fmt.Errorf("config is not defined")
	}
	if config.Channel == "" {
		return nil, fmt.Errorf("channel is not defined")
	}
	if config.ClientOptions == nil {
		config.ClientOptions = &redis.Options{}
	}
	return &Redis{
		ctx:    ctx,
		config: config,
		client: redis.NewClient(config.ClientOptions),
		out:    make(chan interface{}),
	}, nil
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
