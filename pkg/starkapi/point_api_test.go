package starkapi

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testGetPointTypesApiURL = "/core/lists/pointType"
)

func TestPointsUrl(t *testing.T) {
	host := "https://example.com"
	expectedURL := "https://example.com/core/points"

	result := pointsUrl(host)

	assert.Equal(t, expectedURL, result, "Generated URL does not match the expected URL")
}

func TestPointUrl(t *testing.T) {
	host := "https://example.com"
	id := uint32(123)
	expectedURL := "https://example.com/core/points/123"

	result := pointUrl(host, id)

	assert.Equal(t, expectedURL, result, "Generated URL does not match the expected URL")
}

func TestGetAllPointTypes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("%s", testGetPointTypesApiURL) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	api := Client{}
	host := server.URL

	api.Init(host)

	pointApi := api.PointApi

	_, pointErr := pointApi.GetAllPointTypes()
	assert.Equal(t, nil, pointErr)

	badApi := Client{}
	badApi.Init("/bad")

	badPointApi := badApi.PointApi
	_, badFormsApiErr := badPointApi.GetAllPointTypes()
	assert.NotEqual(t, nil, badFormsApiErr)
}
