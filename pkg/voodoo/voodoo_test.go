package voodoo

import (
	"context"
	"fmt"
	"testing"
	"time"

	sChannel "github.com/saromanov/voodoo/pkg/source/channel"
	"github.com/stretchr/testify/assert"
)

func sourceChannel() <-chan interface{} {
	result := make(chan interface{})
	time.Sleep(2 * time.Second)
	go func() {
		for i := 0; i < 25; i++ {
			result <- fmt.Sprintf("Channel source: %v", i)
		}
	}()
	return result
}

func TestAddSource(t *testing.T) {
	n := New()
	assert.Error(t, errNoSources, n.AddSources().Do())

	source, err := sChannel.New(context.TODO(), &sChannel.Options{
		Method: sourceChannel,
	})
	assert.NoError(t, err)
	n = New()
	assert.Error(t, errNoReceivers, n.AddSources(source).Do())
}
