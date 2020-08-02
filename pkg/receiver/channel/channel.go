package channel

import (
	"context"

	"github.com/saromanov/voodoo/pkg/receiver"
)

// Channel defines channel receiver
type Channel struct {
	ctx context.Context
	in  chan interface{}
	f   func(interface{})
}

// New creates channel receiver
func New(ctx context.Context, f func(interface{})) (receiver.Receiver, error) {
	c := &Channel{
		in: make(chan interface{}),
		f:  f,
	}
	go c.init()
	return c, nil
}

// init provides initialization of the main loop
func (r *Channel) init() {
	for msg := range r.in {
		r.f(msg)
	}

	close(r.in)
}

// In provides sending data to the channel
func (r *Channel) In(data interface{}) {
	r.in <- data
}
