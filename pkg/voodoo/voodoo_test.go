package voodoo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSource(t *testing.T) {
	n := New()
	assert.Error(t, errNoSources, n.AddSources().Do())
}
