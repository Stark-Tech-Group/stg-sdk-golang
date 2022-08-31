package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRef_BeginsWith(t *testing.T) {
	a := Ref{Value: "a.xxxxxxxx-xxxxxxxx"}
	assert.True(t, a.BeginsWith("a"))
}

func TestRef_GetPrefix(t *testing.T) {
	a := Ref{Value: "a.xxxxxxxx-xxxxxxxx"}
	assert.Equal(t, "a", a.GetPrefix())

	abc := Ref{Value: "abc.xxxxxxxx-xxxxxxxx"}
	assert.Equal(t, "abc", abc.GetPrefix())
}

func TestRef_String(t *testing.T) {
	a := Ref{Value: "a.xxxxxxxx-xxxxxxxx"}
	assert.Equal(t, "a.xxxxxxxx-xxxxxxxx", fmt.Sprintf("%s", a))
}

func TestCreateRefWithZeroPrefix(t *testing.T) {
	a, error := RandomWithPrefix("")
	assert.NotEqual(t, error, "invalid prefix length")
	assert.NotEqual(t, len(EquipRefType) + len(DefaultPrefixChar) + DefaultLength + len(DefaultSpacingChar), len(a))
}

func TestCreateRefWithLongPrefix(t *testing.T) {
	a, error := RandomWithPrefix("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz")
	assert.Equal(t, error, "invalid prefix length")
	assert.NotEqual(t, len(EquipRefType) + len(DefaultPrefixChar) + DefaultLength + len(DefaultSpacingChar), len(a))
}

func TestCreateRefWithPrefix(t *testing.T) {
	a, error := RandomWithPrefix(EquipRefType)
	assert.Equal(t, error, nil)
	assert.Equal(t, len(EquipRefType) + len(DefaultPrefixChar) + DefaultLength + len(DefaultSpacingChar), len(a))
}

func TestCreateRefWithoutPrefix(t *testing.T) {
	a := Ref{Value: RandomWithoutPrefix()}
	assert.Equal(t, DefaultLength + len(DefaultSpacingChar), len(a.Value))
}
