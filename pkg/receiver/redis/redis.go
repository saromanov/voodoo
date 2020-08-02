package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/saromanov/voodoo/pkg/receiver"
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
func New(ctx context.Context, config *Options) (receiver.Receiver, error) {
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
		client: client,
		in:     make(chan interface{}),
	}
	return r, nil
}

// init provides initialization of the main loop
func (r *Redis) init() {
	defer r.client.Close()
	for msg := range r.in {
		switch m := msg.(type) {
		case string:
			err := r.client.Publish(r.config.Channel, m).Err()
			if err != nil {
				log.Printf("Publish failed with: %s", err)
			}
		default:
			log.Printf("Unsupported redis message type %v", m)
		}
	}
	close(r.in)

	log.Printf("Closing redis producer")
}

// In provides sending data to the channel
func (r *Redis) In(data interface{}) {
	r.in <- data
}
