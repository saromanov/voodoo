package voodoo

import (
	"fmt"
	"sync"
	"time"

	"github.com/saromanov/voodoo/pkg/receiver"
	"github.com/saromanov/voodoo/pkg/source"
	"github.com/saromanov/voodoo/pkg/transform"
)

// Voodoo defines main app
type Voodoo struct {
	sources   []source.Source
	transform transform.Transform
}

// New provides initialization of the app
func New() *Voodoo {
	return &Voodoo{
		sources: []source.Source{},
	}
}

// AddSources provides adding of sources
func (v *Voodoo) AddSources(sources ...source.Source) *Voodoo {
	v.sources = sources
	return v
}

// Transform adds transformation to sources
func (v *Voodoo) Transform(t transform.Transform) *Voodoo {
	v.transform = t
	return v
}

// AddReceivers provides adding of receivers
func (v *Voodoo) AddReceivers(receivers ...receiver.Receiver) *Voodoo {
	return v
}

// Do provides running of the flow
func (v *Voodoo) Do() {
	s := v.sources[0]
	go func() {
		for elem := range s.Out() {
			fmt.Println(elem)
		}
		//close(inlet.In())
	}()
	fmt.Println("STARTED")
	time.Sleep(50 * time.Second)
}

func mergeChannels(cs ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan interface{}) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
