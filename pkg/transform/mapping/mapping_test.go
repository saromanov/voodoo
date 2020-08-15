package mapping

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mapping(d interface{}) interface{} {
	return strings.ToUpper(d.(string))
}
func TestMapping(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()
	r := New(mapping)
	go func() {
		r.In("msg")
	}()

	select {
	case <-ctx.Done():
		t.Fail()
	case res := <-r.Out():
		assert.Equal(t, res, "MSG")
	}
}
