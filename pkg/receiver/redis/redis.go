package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/saromanov/voodoo/pkg/source"
)

// Options defines initialization options for redis receiver
type Options struct {
	ClientOptions *redis.Options
	Channel       string
}

type Redis struct {
	ctx    context.Context
	client *redis.Client
	config *Options
	in     chan interface{}
}

// New creates redis receiver
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
	r := &Redis{
		ctx:    ctx,
		config: config,
		client: redis.NewClient(config.ClientOptions),
		out:    make(chan interface{}),
	}
	return r, nil
}

// Out returns output channel
func (r *Redis) Out() <-chan interface{} {
	return r.out
}
