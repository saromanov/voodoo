package channel

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func sourceChannel() <-chan interface{} {
	result := make(chan interface{})
	time.Sleep(2 * time.Second)
	go func() {
		result <- "foobar"
	}()
	return result
}

func TestChannel(t *testing.T) {

	done := make(chan interface{})
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	go func() {
		r, err := New(context.Background(), &Options{
			Method: sourceChannel,
		})
		assert.NoError(t, err)
		res := <-r.Out()
		done <- res
	}()

	select {
	case res := <-done:
		fmt.Println(&res)
		if res != "foobar" {
			t.Fail()
		}
		return
	case <-ctx.Done():
		assert.Fail(t, "deadline exceed")
	}
}
