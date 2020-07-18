package source

import (
	"github.com/saromanov/voodoo/pkg/transform"
)

// Source is interface for defining source
type Source interface {
	With(transform.Transform) Source
	To() error
}
