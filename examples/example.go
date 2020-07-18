package main

import (
	"context"
	"strings"

	"github.com/saromanov/voodoo/pkg/source/redis"
	"github.com/saromanov/voodoo/pkg/transform/mapping"
)

func mapTransform(data interface{}) interface{} {
	return strings.ToUpper(data.(string))
}

func main() {
	source := redis.New(context.TODO(), nil)
	trans := mapping.New(mapTransform)
	source.With(trans)
}
