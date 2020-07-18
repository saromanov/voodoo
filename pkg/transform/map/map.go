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
	return &Map {
		F:f,
		in:in,
		out: out, 
	}
}