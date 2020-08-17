package voodoo

import (
	"errors"
	"fmt"
	"sync"

	"github.com/saromanov/voodoo/pkg/receiver"
	"github.com/saromanov/voodoo/pkg/source"
	"github.com/saromanov/voodoo/pkg/transform"
)

var (
	errNoSources   = errors.New("sources is not defined")
	errNoReceivers = errors.New("receivers is not defined")
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
	v.sources = append(v.sources, sources...)
	return v
}

// Transform adds transformation to sources
func (v *Voodoo) Transform(t transform.Transform) *Voodoo {
	v.transform = t
	return v
}

// AddReceivers provides adding of receivers
func (v *Voodoo) AddReceivers(receivers ...receiver.Receiver) *Voodoo {
	v.receivers = append(v.receivers, receivers...)
	return v
}

// Do provides running of the flow
func (v *Voodoo) Do() error {
	if err := v.validate(); err != nil {
		return fmt.Errorf("unable to validate flow: %v", err)
	}
	go func() {
		for elem := range getData(v.sources...) {
			v.transform.In(elem)
			for _, r := range v.receivers {
				res := <-v.transform.Out()
				if res != struct{}{} {
					r.In(res)
				}
			}
		}
	}()

	return nil
}

func (v *Voodoo) validate() error {
	if len(v.sources) == 0 {
		return errNoSources
	}
	if len(v.receivers) == 0 {
		return errNoReceivers
	}
	return nil
}

func getData(cs ...source.Source) <-chan interface{} {
	if len(cs) == 1 {
		return cs[0].Out()
	}
	return mergeChannels(cs...)
}

// mergeChannels provides merging of several channels at once
func mergeChannels(cs ...source.Source) <-chan interface{} {
	out := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(s source.Source) {
			for v := range s.Out() {
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
