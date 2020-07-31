package channel

import (
	"context"
	"fmt"
	"log"

	"github.com/saromanov/voodoo/pkg/receiver"
)

// Channel defines channel receiver
type Channel struct {
	ctx context.Context
	in  chan interface{}
}

// New creates channel receiver
func New(ctx context.Context) (receiver.Receiver, error) {
	c := &Channel{
		in: make(chan interface{}),
	}
	go c.init()
	return c, nil
}

// init provides initialization of the main loop
func (r *Channel) init() {
	for msg := range r.in {
		fmt.Println(msg)
	}

	log.Printf("Closing channel receiver")
}

// In provides sending data to the channel
func (r *Channel) In(data interface{}) {
	r.in <- data
}
