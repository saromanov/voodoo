package channel

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChannel(t *testing.T) {

	done := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	go func() {
		rec := func(d interface{}) {
			if d.(string) == "msg" {
				done <- struct{}{}
			}
		}
		r, err := New(context.Background(), rec)
		assert.NoError(t, err)
		r.In("msg")
	}()

	select {
	case <-done:
		return
	case <-ctx.Done():
		assert.Fail(t, "deadline exceed")
	}
}
