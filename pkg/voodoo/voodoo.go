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
	receivers []receiver.Receiver
}

// New provides initialization of the app
func New() *Voodoo {
	return &Voodoo{
		sources:   []source.Source{},
		receivers: []receiver.Receiver{},
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
	v.receivers = receivers
	return v
}

// Do provides running of the flow
func (v *Voodoo) Do() {
	go func() {
		for elem := range getData(v.sources...) {
			v.transform.In(elem)
			for _, r := range v.receivers {
				r.In(<-v.transform.Out())
			}
		}
		//close(inlet.In())
	}()
	fmt.Println("STARTED")
	time.Sleep(50 * time.Second)
}

func getData(cs ...source.Source) <-chan interface{} {
	if len(cs) == 1 {
		return cs[0].Out()
	}
	return mergeChannels(cs...)
}

// mergeChannels provides merging of several channels at one
func mergeChannels(cs ...source.Source) <-chan interface{} {
	out := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan interface{}) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c.Out())
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
