package voodoo

import (
	"github.com/saromanov/voodoo/pkg/receiver"
	"github.com/saromanov/voodoo/pkg/source"
	"github.com/saromanov/voodoo/pkg/transform"
)

// Voodoo defines main app
type Voodoo struct {
}

// AddSources provides adding of sources
func (v *Voodoo) AddSources(sources ...source.Source) *Voodoo {
	return v
}

// Transform adds transformation to sources
func (v *Voodoo) Transform(t transform.Transform) *Voodoo {
	return v
}

// AddReceivers provides adding of receivers
func (v *Voodoo) AddReceivers(receivers ...receiver.Receiver) *Voodoo {
	return v
}
