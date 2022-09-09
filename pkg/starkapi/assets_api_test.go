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

const testHost =  "https://test.com"
const testAssetsApiURL =  "/core/assets"
const testAssetsApiURLWithRef =  "/core/assets/e.test/tags"
const testAssetRef = "e.test"

func TestAssetsApi_hostUrl(t *testing.T) {
	api := Client{}
	host := testHost
	api.Init(host)
	assetsApi := api.AssetsApi
	assert.Equal(t, testHost, assetsApi.host())
}

func TestAssetsApi_baseUrl(t *testing.T) {
	api := Client{}
	host := testHost
	api.Init(host)
	assetsApi := api.AssetsApi
	assert.Equal(t, fmt.Sprintf("%s%s", testHost, testAssetsApiURL), assetsApi.baseUrl())
}

func TestAssetsApi_tagUrlWithRef(t *testing.T) {
	api := Client{}
	host := testHost
	api.Init(host)
	assetsApi := api.AssetsApi
	assert.Equal(t, fmt.Sprintf("%s%s", testHost, testAssetsApiURLWithRef), assetsApi.tagUrlWithRef(testAssetRef))
}

func TestAddTagToAsset(t *testing.T) {
	const badTagValue = "-1"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != testAssetsApiURLWithRef {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var requestTag domain.Tag
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&requestTag)
		if err != nil || requestTag.Value == badTagValue {
			w.WriteHeader(http.StatusBadRequest)
			return
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

	//Test tag with -1 value to force bad request on POST
	badPostErr := assetsApi.AddNewTag(asset, "Test", badTagValue)
	assert.NotEqual(t, nil, badPostErr)

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
		if r.URL.Path !=  testAssetsApiURLWithRef {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.Method !=  http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
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
	const validTagName = "Test"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("%s/%s", testAssetsApiURLWithRef, validTagName) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusBadRequest)
			return
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

	//Test non-existent tag and error on DELETE request
	noTagExistsErr := assetsApi.DeleteTag(asset, "Test123")
	assert.NotEqual(t, nil, noTagExistsErr)

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