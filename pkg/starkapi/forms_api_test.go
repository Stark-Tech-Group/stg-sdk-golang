package starkapi

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testFormsApiURL =  "/core/forms"
const testFormsControlPrefix =  "/controls"

func TestFormsApi_host(t *testing.T) {
	api := Client{}
	host := testHost
	api.Init(host)
	formsApi := api.FormsApi
	assert.Equal(t, testHost, formsApi.host())
}

func TestFormsApi_baseUrl(t *testing.T) {
	api := Client{}
	host := testHost
	api.Init(host)
	formsApi := api.FormsApi
	assert.Equal(t, fmt.Sprintf("%s%s", testHost, testFormsApiURL), formsApi.baseUrl())
}

func TestFormsApi_controlsPrefix(t *testing.T) {
	api := Client{}
	host := testHost
	api.Init(host)
	formsApi := api.FormsApi
	assert.Equal(t, testFormsControlPrefix, formsApi.controlsPrefix())
}

func TestFormsApi_GetAllControls(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("%s%s", testFormsApiURL, testFormsControlPrefix) {
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

	formsApi := api.FormsApi

	_, formsErr := formsApi.GetAllControls()
	assert.Equal(t, nil, formsErr)

	badApi := Client{}
	badApi.Init("/bad")

	badFormsApi := badApi.FormsApi
	_, badFormsApiErr := badFormsApi.GetAllControls()
	assert.NotEqual(t, nil, badFormsApiErr)
}

func TestFormsApi_GetAllControlsForAsset(t *testing.T) {
	const assetRef = "e.test"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("%s/%s%s", testFormsApiURL, assetRef, testFormsControlPrefix) {
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

	formsApi := api.FormsApi

	_, formsErr := formsApi.GetAllControlsForAsset(assetRef)
	assert.Equal(t, nil, formsErr)

	badApi := Client{}
	badApi.Init("/bad")

	badFormsApi := badApi.FormsApi
	_, badFormsApiErr := badFormsApi.GetAllControlsForAsset(assetRef)
	assert.NotEqual(t, nil, badFormsApiErr)
}