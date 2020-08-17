package voodoo

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	rec "github.com/saromanov/voodoo/pkg/receiver/channel"
	sChannel "github.com/saromanov/voodoo/pkg/source/channel"
	"github.com/saromanov/voodoo/pkg/transform/mapping"
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

func mapTransform(data interface{}) interface{} {
	return strings.ToUpper(data.(string))
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

func TestRegister(t *testing.T) {
	n := New()
	source, err := sChannel.New(context.TODO(), &sChannel.Options{
		Method: sourceChannel,
	})

	reca := func(data interface{}) {

	}
	receiver, err := rec.New(context.Background(), reca)
	assert.NoError(t, err)
	assert.NoError(t, err)
	n.AddSources(source).Transform(mapping.New(mapTransform)).AddReceivers(receiver)
}
