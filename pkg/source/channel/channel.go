package channel

import (
	"context"
	"fmt"

	"github.com/saromanov/voodoo/pkg/source"
)

// Options defines initialization options for channel source
type Options struct {
	Method func() <-chan interface{}
}

type Channel struct {
	ctx    context.Context
	config *Options
	out    chan interface{}
	method func() <-chan interface{}
}

// New creates channel source
func New(ctx context.Context, config *Options) (source.Source, error) {
	if config == nil {
		return nil, fmt.Errorf("config is not defined")
	}
	ch := &Channel{
		method: config.Method,
	}
	go ch.init()
	return ch, nil
}

// Out returns output channel
func (r *Channel) Out() <-chan interface{} {
	return r.out
}

// init provides initialization of the receiver from Channel
func (r *Channel) init() {
	defer func() {
		close(r.out)
	}()

	for {
		select {
		case <-r.ctx.Done():
			break
		case msg := <-r.method():
			fmt.Println(msg)
			r.out <- msg
		}
	}
}
