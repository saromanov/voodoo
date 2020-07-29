package main

import (
	"context"
	"strings"

	rec "github.com/saromanov/voodoo/pkg/receiver/redis"
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

	receiver, err := rec.New(context.Background(), &rec.Options{})
	if err != nil {
		panic(err)
	}
	v.AddSources(source).Transform(mapping.New(mapTransform)).AddReceivers(receiver).Do()
}
