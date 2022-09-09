package starkapi

import (
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testHost =  "https://test.com"
const testURLWithRef =  "/core/assets/e.test/tags"

func TestAssetsApi_HostUrl(t *testing.T) {
	api := Client{}
	host := testHost
	api.Init(host)
	assetsApi := api.AssetsApi
	assert.Equal(t, testHost, assetsApi.host())
}

func TestAddTagToAsset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path !=  testURLWithRef {
			t.Errorf("Expected to request '%s', got: %s", testURLWithRef , r.URL.Path)
		}
		if r.Method !=  http.MethodPost {
			t.Errorf("Expected a %s request , got: %s",http.MethodPost, r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	api := Client{}
	host := server.URL

	api.Init(host)

	//Test valid asset
	asset := domain.Asset{
		Id:   1,
		Ref:  "e.test",
		Name: "Test",
		Type: "Equip",
	}

	assetsApi := api.AssetsApi

	assetErr := assetsApi.AddNewTag(asset, "Test", "1")
	assert.Equal(t, nil, assetErr)


	//Test invalid tag
	badTagErr := assetsApi.AddNewTag(asset, "", "1")
	assert.NotEqual(t, nil, badTagErr)
	
	//Test asset with no Ref
	badAsset := domain.Asset{
		Id:   1,
		Ref:  "",
		Name: "",
		Type: "Equip",
	}

	badAssetErr := assetsApi.AddNewTag(badAsset, "Test", "1")
	assert.NotEqual(t, nil, badAssetErr)
}

func TestAddTagsToAsset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path !=  testURLWithRef {
			t.Errorf("Expected to request '%s', got: %s", testURLWithRef , r.URL.Path)
		}
		if r.Method !=  http.MethodPost {
			t.Errorf("Expected a %s request , got: %s",http.MethodPost, r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	api := Client{}
	host := server.URL

	api.Init(host)

	//Test valid asset and tag array
	asset := domain.Asset{
		Id:   1,
		Ref:  "e.test",
		Name: "Test",
		Type: "Equip",
	}

	var tags []domain.Tag
	tags = append(tags, domain.Tag{Name: "Test", Value: "1"})

	assetsApi := api.AssetsApi

	assetErr := assetsApi.AddNewTags(asset, tags)
	assert.Equal(t, nil, assetErr)

	//Test Tag with no name
	tags = append(tags, domain.Tag{Name: "", Value: "1"})
	noNameTagErr := assetsApi.AddNewTags(asset, tags)
	assert.NotEqual(t, nil, noNameTagErr)

	//Test empty tag array
	var emptyTags []domain.Tag
	badTagErr := assetsApi.AddNewTags(asset, emptyTags)
	assert.NotEqual(t, nil, badTagErr)

	//Test asset with no Ref
	badAsset := domain.Asset{
		Id:   1,
		Ref:  "",
		Name: "",
		Type: "Equip",
	}

	badAssetErr := assetsApi.AddNewTag(badAsset, "Test", "1")
	assert.NotEqual(t, nil, badAssetErr)
}

func TestDeleteTagFromAsset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path !=  testURLWithRef {
			t.Errorf("Expected to request '%s', got: %s", testURLWithRef , r.URL.Path)
		}
		if r.Method !=  http.MethodDelete {
			t.Errorf("Expected a %s request , got: %s",http.MethodDelete, r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	api := Client{}
	host := server.URL

	api.Init(host)

	//Test valid asset
	asset := domain.Asset{
		Id:   1,
		Ref:  "e.test",
		Name: "Test",
		Type: "Equip",
	}

	assetsApi := api.AssetsApi

	assetErr := assetsApi.DeleteTag(asset, "Test")
	assert.Equal(t, nil, assetErr)


	//Test invalid tag
	badTagErr := assetsApi.DeleteTag(asset, "")
	assert.NotEqual(t, nil, badTagErr)

	//Test asset with no Ref
	badAsset := domain.Asset{
		Id:   1,
		Ref:  "",
		Name: "",
		Type: "Equip",
	}

	badAssetErr := assetsApi.DeleteTag(badAsset, "Test")
	assert.NotEqual(t, nil, badAssetErr)
}