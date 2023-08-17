package starkapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointsUrl(t *testing.T) {
	host := "http://example.com"
	expectedURL := "http://example.com/core/points"

	result := pointsUrl(host)

	assert.Equal(t, expectedURL, result, "Generated URL does not match the expected URL")
}

func TestPointUrl(t *testing.T) {
	host := "http://example.com"
	id := uint32(123)
	expectedURL := "http://example.com/core/points/123"

	result := pointUrl(host, id)

	assert.Equal(t, expectedURL, result, "Generated URL does not match the expected URL")
}
