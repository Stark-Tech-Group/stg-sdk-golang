package starkapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testFormsApiURL        = "/core/forms"
	testFormsControlPrefix = "/controls"
	testFormName           = "SampleName"
	testFormNameInvalid    = "DoesNotExist"
	errorGetAllControls    = "bad request"
	testFormsControlRef    = "testRef"
	testFormsControlName   = "Test Name"
	testFormsControlString = "Test Control"
)

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

func TestFormsApi_GetControlByName(t *testing.T) {
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
	formsApi.client = &MockClient{
		getFunc: func(url string) ([]byte, error) {
			controlList := domain.FormControlList{
				FormControlList: []*domain.FormControl{
					{Name: "Name1"},
					{Name: "SampleName"},
					{Name: "Name2"},
				},
			}
			data, _ := json.Marshal(controlList)
			return data, nil
		},
		getHostFunc: func() string {
			return "someHost"
		},
	}

	result, err := formsApi.GetControlByName(testFormName)

	// Assert
	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, testFormName, result.Name)

	result, err = formsApi.GetControlByName(testFormNameInvalid)

	// Assert
	assert.NoError(t, err, "Unexpected error")
	assert.NotEqual(t, testFormNameInvalid, result.Name)
	assert.Equal(t, "", result.Name)
}

func TestFormsApi_GetControlByNameErrorInGetAllControls(t *testing.T) {
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
	formsApi.client = &MockClient{
		getFunc: func(url string) ([]byte, error) {
			return nil, errors.New(errorGetAllControls)
		},
		getHostFunc: func() string {
			return "someHost"
		},
	}

	_, err := formsApi.GetControlByName(testFormName)

	// Assert
	assert.Error(t, err, errorGetAllControls)
}

func TestFormsApi_CreateControlOnRef(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("%s/%s%s", testFormsApiURL, testFormsControlRef, testFormsControlPrefix) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.Method != http.MethodPost {
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

	control := domain.FormControl{Ref: testFormsControlRef, Name: testFormsControlName, Control: testFormsControlString}

	res, formsErr := formsApi.CreateControlOnRef(control)
	assert.Equal(t, nil, formsErr)
	assert.Equal(t, control.Ref, res.Ref)
	assert.Equal(t, control.Name, res.Name)
	assert.Equal(t, control.Control, res.Control)

	badApi := Client{}
	badApi.Init("/bad")

	badFormsApi := badApi.FormsApi
	_, badFormsApiErr := badFormsApi.CreateControlOnRef(control)
	assert.NotEqual(t, nil, badFormsApiErr)
}

func TestFormsApi_CreateControlOnRefWithMissingFields(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("%s/%s%s", testFormsApiURL, testFormsControlRef, testFormsControlPrefix) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.Method != http.MethodPost {
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

	controlMissingRef := domain.FormControl{Name: testFormsControlName, Control: testFormsControlString}
	controlMissingName := domain.FormControl{Ref: testFormsControlRef, Control: testFormsControlString}
	controlMissingControl := domain.FormControl{Ref: testFormsControlRef, Name: testFormsControlName}

	_, formsErr := formsApi.CreateControlOnRef(controlMissingRef)
	assert.NotNil(t, formsErr)

	_, formsErr = formsApi.CreateControlOnRef(controlMissingName)
	assert.NotNil(t, formsErr)

	_, formsErr = formsApi.CreateControlOnRef(controlMissingControl)
	assert.NotNil(t, formsErr)

}
