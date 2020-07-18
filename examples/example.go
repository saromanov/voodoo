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
	source, err := redis.New(context.TODO(), &redis.Options{Channel: "test"})
	if err != nil {
		panic(err)
	}
	trans := mapping.New(mapTransform)
	source.With(trans)
}
