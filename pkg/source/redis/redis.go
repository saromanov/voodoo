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
	client := redis.NewClient(config.ClientOptions)
	pubsub := client.Subscribe(config.Channel)
	r := &Redis{
		ctx:    ctx,
		config: config,
		client: redis.NewClient(config.ClientOptions),
		out:    make(chan interface{}),
	}
	go r.init(pubsub.Channel())
	return r, nil
}

// Out returns output channel
func (r *Redis) Out() <-chan interface{} {
	return r.out
}

// prepare converts input redis message to data interface
func (r *Redis) prepare(in interface{}) interface{} {
	result := in.(*redis.Message)
	return result.Payload
}

// init provides initialization of the receiver from redis
func (r *Redis) init(ch <-chan *redis.Message) {
	defer func() {
		close(r.out)
		r.client.Close()
	}()

	for {
		select {
		case <-r.ctx.Done():
			break
		case msg := <-ch:
			r.out <- r.prepare(msg)
		}
	}
}

func (r *Redis) With(transform.Transform) source.Source {
	return r
}

func (r *Redis) To() error {
	return nil
}
