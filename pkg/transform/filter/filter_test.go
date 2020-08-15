package filter

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func filterTransform(data interface{}) bool {
	r := data.(string)
	return len(r) > 2
}

func basicTest(t *testing.T, msg string, expected interface{}) {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()
	r := New(filterTransform)
	go func() {
		r.In(msg)
	}()

	select {
	case <-ctx.Done():
		t.Fail()
	case res := <-r.Out():
		assert.Equal(t, res, expected)
	}
}
func TestFilter(t *testing.T) {
	basicTest(t, "msg", "msg")
	basicTest(t, "m", struct{}{})
}
