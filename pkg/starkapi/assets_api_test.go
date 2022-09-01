package starkapi

import (
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddTagToAsset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path !=  "/core/assets/e.test/tags" {
			t.Errorf("Expected to request '%s', got: %s", "/tags" , r.URL.Path)
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
		Url:  "test/url",
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
		Url:  "test/url",
		Name: "",
		Type: "Equip",
	}

	badAssetErr := assetsApi.AddNewTag(badAsset, "Test", "1")
	assert.NotEqual(t, nil, badAssetErr)
}
