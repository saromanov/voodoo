package filter

import (
	"github.com/saromanov/voodoo/pkg/transform"
)

// Func defines function for transformation
type Func func(interface{}) bool

type Filter struct {
	F   Func
	in  chan interface{}
	out chan interface{}
}

// New initialize method for filter
func New(f Func) transform.Transform {
	r := &Filter{
		F:   f,
		in:  make(chan interface{}),
		out: make(chan interface{}),
	}
	go r.apply()
	return r
}

// In receive data for processing
func (m *Filter) In(elem interface{}) {
	m.in <- elem
}

// Out returns processed data
func (m *Filter) Out() <-chan interface{} {
	return m.out
}

// apply provides doing of map data
func (m *Filter) apply() {
	for elem := range m.in {
		go func(e interface{}) {
			res := m.F(e)
			if res {
				m.out <- e
			} else {
				m.out <- struct{}{}
			}
		}(elem)
	}
	close(m.out)
}
