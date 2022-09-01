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

func TestCreateRefWithZeroPrefix(t *testing.T) {
	a, err := domain.RandomWithPrefix("")
	assert.NotEqual(t, err, "invalid prefix length")
	assert.NotEqual(t, len(domain.EquipRefType) + len(domain.defaultPrefixChar) + domain.defaultLength + len(domain.defaultSpacingChar), len(a))
}

func TestCreateRefWithLongPrefix(t *testing.T) {
	a, err := domain.RandomWithPrefix("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz")
	assert.Equal(t, err, "invalid prefix length")
	assert.NotEqual(t, len(domain.EquipRefType) + len(domain.defaultPrefixChar) + domain.defaultLength + len(domain.defaultSpacingChar), len(a))
}

func TestCreateRefWithPrefix(t *testing.T) {
	a, err := domain.RandomWithPrefix(domain.EquipRefType)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(domain.EquipRefType) + len(domain.defaultPrefixChar) + domain.defaultLength + len(domain.defaultSpacingChar), len(a))
}

func TestCreateRefWithoutPrefix(t *testing.T) {
	a := domain.Ref{Value: domain.RandomWithoutPrefix()}
	assert.Equal(t, domain.defaultLength + len(domain.defaultSpacingChar), len(a.Value))
}
