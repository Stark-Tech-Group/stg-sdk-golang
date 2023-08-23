package starkapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoginUrl(t *testing.T) {
	host := "https://example.com"
	expectedURL := "https://example.com/login"

	url := loginUrl(host)

	assert.Equal(t, expectedURL, url)
}

func TestMeUrl(t *testing.T) {
	host := "https://example.com"
	expectedURL := "https://example.com/core/persons/me"

	url := meUrl(host)

	assert.Equal(t, expectedURL, url)
}
