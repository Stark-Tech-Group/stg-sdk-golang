package domain

import (
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRef_BeginsWith(t *testing.T) {
	a := domain.Ref{Value: "a.xxxxxxxx-xxxxxxxx"}
	assert.True(t, a.BeginsWith("a"))
}

func TestRef_GetPrefix(t *testing.T) {
	a := domain.Ref{Value: "a.xxxxxxxx-xxxxxxxx"}
	assert.Equal(t, "a", a.GetPrefix())

	abc := domain.Ref{Value: "abc.xxxxxxxx-xxxxxxxx"}
	assert.Equal(t, "abc", abc.GetPrefix())
}

func TestRef_String(t *testing.T) {
	a := domain.Ref{Value: "a.xxxxxxxx-xxxxxxxx"}
	assert.Equal(t, "a.xxxxxxxx-xxxxxxxx", fmt.Sprintf("%s", a))
}
