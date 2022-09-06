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

	noPrefix := Ref{Value: "xxxxxxxx-xxxxxxxx"}
	assert.Equal(t, "", noPrefix.GetPrefix())
}

func TestRef_String(t *testing.T) {
	a := Ref{Value: "a.xxxxxxxx-xxxxxxxx"}
	assert.Equal(t, "a.xxxxxxxx-xxxxxxxx", fmt.Sprintf("%s", a))
}

func TestCreateRefWithZeroPrefix(t *testing.T) {
	a, err := RandomWithPrefix("")
	assert.NotEqual(t, err, "invalid prefix length")
	assert.NotEqual(t, len(EquipRefType)+len(defaultPrefixChar)+defaultLength+len(defaultSpacingChar), len(a))
}

func TestCreateRefWithLongPrefix(t *testing.T) {
	a, err := RandomWithPrefix("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz")
	assert.Equal(t, err.Error(), "invalid prefix length")
	assert.NotEqual(t, len(EquipRefType)+len(defaultPrefixChar)+defaultLength+len(defaultSpacingChar), len(a))
}

func TestCreateRefWithPrefix(t *testing.T) {
	a, err := RandomWithPrefix(EquipRefType)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(EquipRefType)+len(defaultPrefixChar)+defaultLength+len(defaultSpacingChar), len(a))
}

func TestCreateRefWithoutPrefix(t *testing.T) {
	a := Ref{Value: RandomWithoutPrefix()}
	assert.Equal(t, defaultLength+len(defaultSpacingChar), len(a.Value))
}
