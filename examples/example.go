package main

import (
	"context"
	"fmt"
	"strings"

	rec "github.com/saromanov/voodoo/pkg/receiver/channel"
	"github.com/saromanov/voodoo/pkg/source/redis"
	"github.com/saromanov/voodoo/pkg/transform/mapping"
	"github.com/saromanov/voodoo/pkg/voodoo"
)

func mapTransform(data interface{}) interface{} {
	return strings.ToUpper(data.(string))
}

func reca(data interface{}) {
	fmt.Println(data)
}

func main() {
	v := voodoo.New()
	source, err := redis.New(context.TODO(), &redis.Options{Channel: "test"})
	if err != nil {
		panic(err)
	}

	source2, err := redis.New(context.TODO(), &redis.Options{Channel: "test2"})
	if err != nil {
		panic(err)
	}

	receiver, err := rec.New(context.Background(), reca)
	if err != nil {
		panic(err)
	}
	v.AddSources(source).
		AddSources(source2).
		Transform(mapping.New(mapTransform)).
		AddReceivers(receiver).Do()
	select {}
}
