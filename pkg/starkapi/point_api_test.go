package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testGetPointTypesApiURL = "/core/lists/pointType"
	testPointTypeName1      = "Name1"
	testPointTypeId1        = int32(1)
	testPointTypeName2      = "Name2"
	testPointTypeId2        = int32(2)
	testPointTypeName3      = "Name3"
	testPointTypeId3        = int32(3)
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

func TestGetAllPointTypesWithResp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("%s", testGetPointTypesApiURL) {
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

	pointApi := api.PointApi
	pointApi.client = &MockClient{
		getFunc: func(url string) ([]byte, error) {
			pointTypeList := domain.PointTypes{
				Points: []domain.PointType{
					{Name: "Name1", Id: 1},
					{Name: "Name2", Id: 2},
					{Name: "Name3", Id: 3},
				},
			}
			data, _ := json.Marshal(pointTypeList)
			return data, nil
		},
		getHostFunc: func() string {
			return "someHost"
		},
		postFunc: func(url string, body []byte) ([]byte, error) {
			return nil, nil
		},
	}

	res, pointTypesErr := pointApi.GetAllPointTypes()
	assert.Equal(t, nil, pointTypesErr)
	assert.Equal(t, testPointTypeName1, res.Points[0].Name)
	assert.Equal(t, testPointTypeId1, res.Points[0].Id)
	assert.Equal(t, testPointTypeName2, res.Points[1].Name)
	assert.Equal(t, testPointTypeId2, res.Points[1].Id)
	assert.Equal(t, testPointTypeName3, res.Points[2].Name)
	assert.Equal(t, testPointTypeId3, res.Points[2].Id)

}
