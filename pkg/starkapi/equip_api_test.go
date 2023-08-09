package starkapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquipsUrl(t *testing.T) {
	host := "https://example.com"
	expectedURL := "https://example.com/core/equips"

	url := equipsUrl(host)

	assert.Equal(t, expectedURL, url)
}

func TestEquipUrl(t *testing.T) {
	host := "https://example.com"
	id := uint32(123)
	expectedURL := "https://example.com/core/equips/123"

	url := equipUrl(host, id)

	assert.Equal(t, expectedURL, url)
}
