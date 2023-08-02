package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testFormsApiURL = "/core/forms"
const testFormsControlPrefix = "/controls"

type MockFormsApi struct {
	mock.Mock
}

func (m *MockFormsApi) GetAllControls() (domain.FormControlList, error) {
	args := m.Called()
	return args.Get(0).(domain.FormControlList), args.Error(1)
}

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
	//mockFormsApi := new(MockFormsApi)

	// Define sample data
	name := "SampleName"
	//controlList := domain.FormControlList{
	//	FormControlList: []*domain.FormControl{
	//		{Name: "Name1"},
	//		{Name: "SampleName"},
	//		{Name: "Name2"},
	//	},
	//}

	// Set up the expected behavior of the mock
	//mockFormsApi.On("GetControlByName").On("GetAllControls").Return(controlList, nil)

	// Act
	result, err := formsApi.GetControlByName(name)

	// Assert
	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, name, result.Name, "Control name doesn't match")
}
