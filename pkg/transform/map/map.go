package map

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/saromanov/voodoo/pkg/transform"
)

// Func defines function for transformation
type Func func(interface{}) interface{}

type Map struct {
	F        Func
	in          chan interface{}
	out         chan interface{}
}

// New initialize method for map
func New(f Func)transform.Transform {
	r := &Map {
		F:f,
		in:in,
		out: out, 
	}
	go r.apply()
	return r
}

func (m *Map) With(t transform.Transform) transform.Transform {
	return m
}

// apply provides doing of map data
func (m *Map) apply() {
	for elem := range m.in {
		go func(e interface{}) {
			m.out <- m.F(e)
		}(elem)
	}
	close(m.out)
}