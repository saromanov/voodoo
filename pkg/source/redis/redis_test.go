package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func publishChannel(t *testing.T, channelName string) {
	client := redis.NewClient(&redis.Options{})
	err := client.Publish(channelName, "foobar").Err()
	assert.NoError(t, err)
}

func TestChannel(t *testing.T) {

	channelName := "test-source"
	done := make(chan interface{})
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	go func() {
		r, err := New(context.Background(), &Options{Channel: channelName})
		assert.NoError(t, err)
		res := <-r.Out()
		done <- res
	}()

	time.Sleep(1 * time.Second)
	publishChannel(t, channelName)
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

func TestChannelWithErrors(t *testing.T) {
	_, err := New(context.Background(), nil)
	assert.Error(t, err)

	_, err = New(context.Background(), &Options{})
	assert.Error(t, err)

}
