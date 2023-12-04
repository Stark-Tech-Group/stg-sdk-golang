package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
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

func TestRandomWithoutPrefix_Over1MillionAttempts(t *testing.T) {
	generatedStrings := make(map[string]struct{})

	const numAttempts = 1000000

	for i := 0; i < numAttempts; i++ {
		s := RandomWithoutPrefix()

		if _, exists := generatedStrings[s]; exists {
			t.Fatalf("expected more random but found duplicate: %s", s)
		}

		generatedStrings[s] = struct{}{}
	}
}

func TestNewRef(t *testing.T) {
	prefix := "Example"
	ref := NewRef(prefix)

	if !strings.HasPrefix(ref, prefix) {
		t.Errorf("Expected reference to have prefix %s, but got %s", prefix, ref)
	}

	longPrefix := "ThisIsAVeryLongPrefixThatExceedsTheMaxLength"
	expectedPrefix := longPrefix[:10]
	refWithLongPrefix := NewRef(longPrefix)

	if !strings.HasPrefix(refWithLongPrefix, expectedPrefix) {
		t.Errorf("Expected reference to have truncated prefix %s, but got %s", expectedPrefix, refWithLongPrefix)
	}
}

func TestGetAssetType(t *testing.T) {
	assert.Equal(t, "Site", GetAssetType("s.xxxxxxxx-xxxxxxxx"))
	assert.Equal(t, "Equip", GetAssetType("e.xxxxxxxx-xxxxxxxx"))
	assert.Equal(t, "Point", GetAssetType("p.xxxxxxxx-xxxxxxxx"))
	assert.Equal(t, "Profile", GetAssetType("r.xxxxxxxx-xxxxxxxx"))
	assert.Equal(t, "Form Control Ref", GetAssetType("k.xxxxxxxx-xxxxxxxx"))
	assert.Equal(t, "", GetAssetType("xxxxxxxx-xxxxxxxx"))
}
