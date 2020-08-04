package mapping

import (
	"github.com/saromanov/voodoo/pkg/transform"
)

// Func defines function for transformation
type Func func(interface{}) interface{}

type Mapping struct {
	F   Func
	in  chan interface{}
	out chan interface{}
}

// New initialize method for map
func New(f Func) transform.Transform {
	r := &Mapping{
		F:   f,
		in:  make(chan interface{}),
		out: make(chan interface{}),
	}
	go r.apply()
	return r
}

// In receive data for processing
func (m *Mapping) In(elem interface{}) {
	m.in <- elem
}

// Out returns processed data
func (m *Mapping) Out() <-chan interface{} {
	return m.out
}

// apply provides doing of map data
func (m *Mapping) apply() {
	for elem := range m.in {
		go func(e interface{}) {
			m.out <- m.F(e)
		}(elem)
	}
	close(m.out)
}
