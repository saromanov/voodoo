package main

import (
	"context"
	"fmt"
	"time"

	rec "github.com/saromanov/voodoo/pkg/receiver/channel"
	sChannel "github.com/saromanov/voodoo/pkg/source/channel"
	"github.com/saromanov/voodoo/pkg/source/redis"
	"github.com/saromanov/voodoo/pkg/transform/filter"
	"github.com/saromanov/voodoo/pkg/voodoo"
)

func filterTransform(data interface{}) bool {
	r := data.(string)
	return len(r) > 8
}

func sourceChannel() <-chan interface{} {
	result := make(chan interface{})
	time.Sleep(2 * time.Second)
	go func() {
		for i := 0; i < 25; i++ {
			result <- fmt.Sprintf("%v", i)
		}
	}()
	return result
}
func reca(data interface{}) {
	fmt.Println("RECEIVE: ", data)
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

	source3, err := sChannel.New(context.TODO(), &sChannel.Options{
		Method: sourceChannel,
	})

	if err != nil {
		panic(err)
	}

	receiver, err := rec.New(context.Background(), reca)
	if err != nil {
		panic(err)
	}
	v.AddSources(source).
		AddSources(source2).
		AddSources(source3).
		Transform(filter.New(filterTransform)).
		AddReceivers(receiver).Do()
	select {}
}
