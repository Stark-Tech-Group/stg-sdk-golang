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
	testFormsApiURL           = "/core/forms"
	testFormsControlPrefix    = "/controls"
	testFormName              = "SampleName"
	testFormNameInvalid       = "DoesNotExist"
	errorGetAllControls       = "bad request"
	testFormsControlName      = "Sample Name"
	testFormControlRef        = "j.1111.2222"
	testIssueTargetRef        = "testTargetRef"
	testFormControlValue      = "test value"
	testFormControlDesc       = "test description"
	testErrorBadPost          = "bad post error"
	testErrorGetControlByName = "get control by name error"
	errInvalidFormControl     = "invalid form control provided with name [%s]"
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
		if r.URL.Path != fmt.Sprintf("%s/%s%s", testFormsApiURL, testIssueTargetRef, testFormsControlPrefix) {
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
	formsApi.client = &MockClient{
		getFunc: func(url string) ([]byte, error) {
			controlList := domain.FormControlList{
				FormControlList: []*domain.FormControl{
					{Name: "Name1"},
					getValidFormControl(),
					{Name: "Name2"},
				},
			}
			data, _ := json.Marshal(controlList)
			return data, nil
		},
		getHostFunc: func() string {
			return "someHost"
		},
		postFunc: func(url string, body []byte) ([]byte, error) {
			return nil, nil
		},
	}

	res, formsErr := formsApi.CreateControlOnRef(testFormsControlName, testIssueTargetRef, testFormControlValue)
	assert.Equal(t, nil, formsErr)
	assert.Equal(t, testIssueTargetRef, res.TargetRef)
	assert.Equal(t, testFormsControlName, res.Name)
	assert.Equal(t, testFormControlValue, res.Value)
}

func TestFormsApi_CreateControlOnRefBadReturn(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("%s/%s%s", testFormsApiURL, testIssueTargetRef, testFormsControlPrefix) {
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
	formsApi.client = &MockClient{
		getFunc: func(url string) ([]byte, error) {
			controlList := domain.FormControlList{
				FormControlList: []*domain.FormControl{
					{Name: "Name1"},
					getValidFormControl(),
					{Name: "Name2"},
				},
			}
			data, _ := json.Marshal(controlList)
			return data, nil
		},
		getHostFunc: func() string {
			return "someHost"
		},
		postFunc: func(url string, body []byte) ([]byte, error) {
			return []byte("1"), nil
		},
	}

	_, formsErr := formsApi.CreateControlOnRef(testFormsControlName, testIssueTargetRef, testFormControlValue)
	var jsonUnmarshalError *json.UnmarshalTypeError
	assert.True(t, errors.As(formsErr, &jsonUnmarshalError))
}

func TestFormsApi_CreateControlOnRefBadGetControl(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("%s/%s%s", testFormsApiURL, testIssueTargetRef, testFormsControlPrefix) {
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
	formsApi.client = &MockClient{
		getFunc: func(url string) ([]byte, error) {
			return nil, errors.New(testErrorGetControlByName)
		},
		getHostFunc: func() string {
			return "someHost"
		},
	}

	_, formsErr := formsApi.CreateControlOnRef(testFormsControlName, testIssueTargetRef, testFormControlValue)
	assert.Equal(t, errors.New(testErrorGetControlByName), formsErr)
}

func TestFormsApi_CreateControlOnRefBadPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("%s/%s%s", testFormsApiURL, testIssueTargetRef, testFormsControlPrefix) {
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
	formsApi.client = &MockClient{
		getFunc: func(url string) ([]byte, error) {
			controlList := domain.FormControlList{
				FormControlList: []*domain.FormControl{
					{Name: "Name1"},
					getValidFormControl(),
					{Name: "Name2"},
				},
			}
			data, _ := json.Marshal(controlList)
			return data, nil
		},
		getHostFunc: func() string {
			return "someHost"
		},
		postFunc: func(url string, body []byte) ([]byte, error) {
			return nil, errors.New(testErrorBadPost)
		},
	}

	_, formsErr := formsApi.CreateControlOnRef(testFormsControlName, testIssueTargetRef, testFormControlValue)
	assert.Equal(t, errors.New(testErrorBadPost), formsErr)
}

func TestFormsApi_CreateControlOnRefInvalidFormControlRet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("%s/%s%s", testFormsApiURL, testIssueTargetRef, testFormsControlPrefix) {
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
	formsApi.client = &MockClient{
		getFunc: func(url string) ([]byte, error) {
			controlList := domain.FormControlList{
				FormControlList: []*domain.FormControl{
					{Name: "Name1"},
					getInvalidFormControl(),
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

	_, formsErr := formsApi.CreateControlOnRef(testFormsControlName, testIssueTargetRef, testFormControlValue)
	assert.Equal(t, errors.New(fmt.Sprintf(errInvalidControlNameProvided, testFormsControlName)), formsErr)
}

func TestFormsApi_CreateControlOnRefBadJSONInControl(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("%s/%s%s", testFormsApiURL, testIssueTargetRef, testFormsControlPrefix) {
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
	formsApi.client = &MockClient{
		getFunc: func(url string) ([]byte, error) {
			controlList := domain.FormControlList{
				FormControlList: []*domain.FormControl{
					{Name: "Name1"},
					getInvalidFormControlWithBadJSON(),
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

	_, formsErr := formsApi.CreateControlOnRef(testFormsControlName, testIssueTargetRef, testFormControlValue)
	var jsonSyntaxError *json.SyntaxError
	assert.True(t, errors.As(formsErr, &jsonSyntaxError))
}

func TestFormsApi_CreateControlOnRefWithMissingFields(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("%s/%s%s", testFormsApiURL, testIssueTargetRef, testFormsControlPrefix) {
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

	_, formsErr := formsApi.CreateControlOnRef(testFormsControlName, testIssueTargetRef, "")
	assert.NotNil(t, formsErr)

	_, formsErr = formsApi.CreateControlOnRef(testFormsControlName, "", testFormControlValue)
	assert.NotNil(t, formsErr)

	_, formsErr = formsApi.CreateControlOnRef("", testIssueTargetRef, testFormControlValue)
	assert.NotNil(t, formsErr)
}

func getValidFormControl() *domain.FormControl {
	return &domain.FormControl{
		Id:          1,
		Name:        testFormsControlName,
		Ref:         testFormControlRef,
		Enabled:     true,
		Description: testFormControlDesc,
		Control:     "{\"key\": \"text\",  \"type\": \"text\",  \"templateOptions\": {    \"label\": \"Text\", \"placeholder\": \"Name, email or phone number of Area Manager\", \"required\": true  }}",
	}
}

func getInvalidFormControl() *domain.FormControl {
	return &domain.FormControl{
		Name:        testFormsControlName,
		Ref:         testFormControlRef,
		Enabled:     true,
		Description: testFormControlDesc,
		Control:     "{\"key\": \"text\",  \"type\": \"text\",  \"templateOptions\": {    \"label\": \"Text\", \"placeholder\": \"Name, email or phone number of Area Manager\", \"required\": true  }}",
	}
}

func getInvalidFormControlWithBadJSON() *domain.FormControl {
	return &domain.FormControl{
		Id:          1,
		Name:        testFormsControlName,
		Ref:         testFormControlRef,
		Enabled:     true,
		Description: testFormControlDesc,
		Control:     "invalid json",
	}
}
