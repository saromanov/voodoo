package main

import (
	"context"
	"strings"

	"github.com/saromanov/voodoo/pkg/source/redis"
	"github.com/saromanov/voodoo/pkg/transform/mapping"
	"github.com/saromanov/voodoo/pkg/voodoo"
)

func mapTransform(data interface{}) interface{} {
	return strings.ToUpper(data.(string))
}

func main() {
	v := voodoo.New()
	source, err := redis.New(context.TODO(), &redis.Options{Channel: "test"})
	if err != nil {
		panic(err)
	}
	v.AddSources(source).Transform(mapping.New(mapTransform)).AddReceivers().Do()
	trans := mapping.New(mapTransform)
	source.With(trans)
}
