package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFormControlRef(t *testing.T) {
	n := NewFormControlRef()
	assert.NotNil(t, n)
	assert.Equal(t, FormControlRefTable, n.Ref[0:1])
}
